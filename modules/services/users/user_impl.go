package services

import (
	"context"
	"fmt"
	"time"

	modelsUser "github.com/Billy278/pos_app_monolic/modules/models/users"
	repository "github.com/Billy278/pos_app_monolic/modules/repository/users"
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
	tNow := time.Now()
	userIn.Updated_At = &tNow
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
func (srv *UserSrvImpl) RepoFindUser(ctx context.Context, username string) (err error) {
	fmt.Println("RepoFindUser")
	err = srv.UserRepo.RepoFindUser(ctx, username)
	return
}
