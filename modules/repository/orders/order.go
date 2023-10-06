package repository

import (
	"context"
	"database/sql"

	modelsOrder "github.com/Billy278/pos_app_monolic/modules/models/orders"
)

type OrderRepo interface {
	RepoDetail(ctx context.Context, db *sql.DB, id uint64) (resOrder modelsOrder.Order, err error)
	RepoListUn(ctx context.Context, db *sql.DB, limit, offset uint64) (resOrder []modelsOrder.Order, err error)
	RepoList(ctx context.Context, db *sql.DB, limit, offset uint64) (resOrder []modelsOrder.Order, err error)
	RepoFindByid(ctx context.Context, db *sql.DB, id uint64) (resOrder modelsOrder.Order, err error)
	RepoCreate(ctx context.Context, tx *sql.Tx, orderIn modelsOrder.Order) (resOrder modelsOrder.Order, err error)
	RepoDelete(ctx context.Context, db *sql.DB, id uint64) (err error)
}
