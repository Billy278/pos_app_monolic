package repository

import (
	"context"

	modelsUser "github.com/Billy278/pos_app_monolic/modules/models/users"
)

type UserRepo interface {
	RepoList(ctx context.Context, limit, offset uint64) (resUser []modelsUser.User, err error)
	RepoFindUser(ctx context.Context, username string) (err error)
	RepoFindByid(ctx context.Context, id uint64) (resUser modelsUser.User, err error)
	RepoCreate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error)
	RepoUpdate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error)
	RepoDelete(ctx context.Context, id uint64) (err error)
	RepoFindUsernameToLogin(ctx context.Context, username string) (resUser modelsUser.User, err error)
}
