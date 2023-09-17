package repository

import (
	"context"

	modelsCategories "github.com/Billy278/pos_app_monolic/modules/models/categories"
)

type Categories interface {
	RepoList(ctx context.Context, limit, offset uint64) (resCategories []modelsCategories.Categories, err error)
	RepoFindByid(ctx context.Context, id uint64) (resCategories modelsCategories.Categories, err error)
	RepoCreate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error)
	RepoUpdate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error)
	RepoDelete(ctx context.Context, id uint64) (err error)
}
