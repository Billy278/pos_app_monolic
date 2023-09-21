package services

import (
	"context"

	modelsProducts "github.com/Billy278/pos_app_monolic/modules/models/products"
)

type SrvProduct interface {
	SrvList(ctx context.Context, limit, offset uint64) (resProduct []modelsProducts.Products, err error)
	SrvFindByid(ctx context.Context, id uint64) (resProduct modelsProducts.Products, err error)
	SrvCreate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error)
	SrvUpdate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error)
	SrvDelete(ctx context.Context, id uint64) (err error)
}
