package services

import (
	"context"

	modelsToken "github.com/Billy278/pos_app_monolic/modules/models/tokens"
	modelsUser "github.com/Billy278/pos_app_monolic/modules/models/users"
)

type UserSrv interface {
	SrvList(ctx context.Context, limit, offset uint64) (resUser []modelsUser.User, err error)
	SrvFindByid(ctx context.Context, id uint64) (resUser modelsUser.User, err error)
	SrvCreate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error)
	SrvUpdate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error)
	SrvDelete(ctx context.Context, id uint64) (err error)
	SrvFindUser(ctx context.Context, username string) (err error)
	SrvFindUsernameToLogin(ctx context.Context, username, password string) (resToken modelsToken.Tokens, err error)
}
