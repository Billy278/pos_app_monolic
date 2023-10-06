package crypto

import (
	"errors"
	"os"

	"github.com/kataras/jwt"
)

func CreatedJWT(claim any) (token string, err error) {
	//create /sign jwt
	key := os.Getenv("Key")
	tokenClaim, err := jwt.Sign(jwt.HS256, []byte(key), claim)
	if err != nil {
		err = errors.New("ERROR SIGN CLAIM")
		return
	}
	token = string(tokenClaim)
	return
}

func ParseAndVerifyToken(token string, claim any) (err error) {
	//verify token
	key := os.Getenv("Key")
	verifyToken, err := jwt.Verify(jwt.HS256, []byte(key), []byte(token))
	if err != nil {
		err = errors.New("ERROR PARSE JWT")
		return
	}
	err = verifyToken.Claims(&claim)
	return
}
