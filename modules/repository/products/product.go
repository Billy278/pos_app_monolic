package repository

import (
	"context"
	"database/sql"

	modelsProducts "github.com/Billy278/pos_app_monolic/modules/models/products"
)

type Product interface {
	RepoList(ctx context.Context, limit, offset uint64) (resProduct []modelsProducts.Products, err error)
	RepoFindByid(ctx context.Context, id uint64) (resProduct modelsProducts.Products, err error)
	RepoCreate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error)
	RepoUpdate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error)
	RepoDelete(ctx context.Context, id uint64) (err error)
	RepoUpdateTx(ctx context.Context, db *sql.Tx, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error)
	RepoFindByidTx(ctx context.Context, db *sql.Tx, id uint64) (resProduct modelsProducts.Products, err error)
}
