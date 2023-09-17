package services

import (
	"context"

	modelsCategories "github.com/Billy278/pos_app_monolic/modules/models/categories"
)

type SrvCategories interface {
	SrvList(ctx context.Context, limit, offset uint64) (resCategories []modelsCategories.Categories, err error)
	SrvFindByid(ctx context.Context, id uint64) (resCategories modelsCategories.Categories, err error)
	SrvCreate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error)
	SrvUpdate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error)
	SrvDelete(ctx context.Context, id uint64) (err error)
}
