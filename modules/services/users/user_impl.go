package services

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	modelsToken "github.com/Billy278/pos_app_monolic/modules/models/tokens"
	modelsUser "github.com/Billy278/pos_app_monolic/modules/models/users"
	repository "github.com/Billy278/pos_app_monolic/modules/repository/users"
	"github.com/Billy278/pos_app_monolic/pkg/crypto"
)

type UserSrvImpl struct {
	UserRepo repository.UserRepo
}

func NewUserSrvImpl(userrepo repository.UserRepo) UserSrv {
	return &UserSrvImpl{
		UserRepo: userrepo,
	}
}

func (srv *UserSrvImpl) SrvList(ctx context.Context, limit, offset uint64) (resUser []modelsUser.User, err error) {
	fmt.Println("SrvList")
	resUser, err = srv.UserRepo.RepoList(ctx, limit, offset)
	if err != nil {
		return
	}

	return
}
func (srv *UserSrvImpl) SrvFindByid(ctx context.Context, id uint64) (resUser modelsUser.User, err error) {
	fmt.Println("SrvFindByid")
	resUser, err = srv.UserRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	return
}
func (srv *UserSrvImpl) SrvCreate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error) {
	fmt.Println("SrvCreate")
	tNow := time.Now()
	userIn.Created_At = &tNow
	hashPass, err := crypto.GenereteHash(userIn.Password)
	if err != nil {
		return
	}
	userIn.Password = hashPass
	resUser, err = srv.UserRepo.RepoCreate(ctx, userIn)
	if err != nil {
		return
	}
	return
}
func (srv *UserSrvImpl) SrvUpdate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error) {
	fmt.Println("SrvUpdate")
	_, err = srv.UserRepo.RepoFindByid(ctx, userIn.Id)
	if err != nil {
		return
	}
	hashPass, err := crypto.GenereteHash(userIn.Password)
	if err != nil {
		return
	}
	tNow := time.Now()
	userIn.Updated_At = &tNow
	userIn.Password = hashPass
	resUser, err = srv.UserRepo.RepoUpdate(ctx, userIn)
	if err != nil {
		return
	}
	return
}
func (srv *UserSrvImpl) SrvDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("SrvDelete")
	_, err = srv.UserRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	err = srv.UserRepo.RepoDelete(ctx, id)
	if err != nil {
		return
	}

	return
}
func (srv *UserSrvImpl) SrvFindUser(ctx context.Context, username string) (err error) {
	fmt.Println("SrvFindUser")
	err = srv.UserRepo.RepoFindUser(ctx, username)
	return
}
func (srv *UserSrvImpl) SrvFindUsernameToLogin(ctx context.Context, username, password string) (resToken modelsToken.Tokens, err error) {
	fmt.Println("SrvFindUsernameToLogin")
	resUser, err := srv.UserRepo.RepoFindUsernameToLogin(ctx, username)
	if err != nil {
		return
	}
	err = crypto.CompareHash(resUser.Password, password)
	if err != nil {
		return
	}
	idToken, accessToken, refreshToken, err := srv.GenerateAllTokenWithConcurency(ctx, resUser.Id, resUser.Name, resUser.Email, "jti")
	if err != nil {
		return
	}
	resToken.IDToken = idToken
	resToken.AccessToken = accessToken
	resToken.RefreshToken = refreshToken
	return
}

func (srv *UserSrvImpl) GenerateAllTokenWithConcurency(ctx context.Context, userId uint64, name, email, jti string) (idToken, accessToken, refreshToken string, err error) {
	fmt.Println("GenerateAllTokenWithConcurency")
	tNow := time.Now()
	defaultClaim := modelsToken.DefaultClaim{
		Expired:   int(tNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(time.Now().Unix()),
		IssuedAt:  int(time.Now().Unix()),
		Issuer:    fmt.Sprint(userId),
		Audience:  "pop_app_monolic",
		JTI:       jti,
		Type:      modelsToken.ID_TOKEN,
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go func(_defaultClaim modelsToken.DefaultClaim) {
		defer wg.Done()
		_defaultClaim.Expired = int(time.Now().Add(2 * time.Hour).Unix())
		_defaultClaim.Type = modelsToken.ACCESS_TOKEN
		//generated acccess claim
		accessTokenClaim := struct {
			modelsToken.DefaultClaim
			modelsToken.AccessClaim
		}{
			_defaultClaim,
			modelsToken.AccessClaim{
				UserId: fmt.Sprint(userId),
				Name:   name,
				Email:  email,
			},
		}
		accessToken, err = crypto.CreatedJWT(accessTokenClaim)
		if err != nil {
			log.Println("Error create access token")
		}
	}(defaultClaim)
	go func(_defaultClaim modelsToken.DefaultClaim) {
		defer wg.Done()
		//generate id claim
		idTokenClaim := struct {
			modelsToken.DefaultClaim
			modelsToken.IdClaim
		}{
			_defaultClaim,
			modelsToken.IdClaim{
				UserId: fmt.Sprint(userId),
				Name:   name,
				Email:  email,
			},
		}
		idToken, err = crypto.CreatedJWT(idTokenClaim)
		if err != nil {
			log.Println("Error create id token")
		}

	}(defaultClaim)

	go func(_defaultClaim modelsToken.DefaultClaim) {
		defer wg.Done()
		//generate refresh token
		_defaultClaim.Expired = int(time.Now().Add(time.Hour).Unix())
		_defaultClaim.Type = modelsToken.REFRESH_TOKEN
		refreshClaimToken := struct {
			modelsToken.DefaultClaim
		}{
			_defaultClaim,
		}
		refreshToken, err = crypto.CreatedJWT(refreshClaimToken)
		if err != nil {
			log.Println("Error create refresh token")
		}

	}(defaultClaim)
	wg.Wait()
	return
}
