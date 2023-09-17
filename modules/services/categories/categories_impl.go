package services

import (
	"context"
	"fmt"
	"time"

	modelsCategories "github.com/Billy278/pos_app_monolic/modules/models/categories"
	repository "github.com/Billy278/pos_app_monolic/modules/repository/categories"
)

type SrvCategoriesImpl struct {
	CategoriesRepo repository.Categories
}

func NewSrvCategoriesImpl(categoriesRepo repository.Categories) SrvCategories {
	return &SrvCategoriesImpl{
		CategoriesRepo: categoriesRepo,
	}
}
func (srv *SrvCategoriesImpl) SrvList(ctx context.Context, limit, offset uint64) (resCategories []modelsCategories.Categories, err error) {
	fmt.Println("SrvList")
	resCategories, err = srv.CategoriesRepo.RepoList(ctx, limit, offset)
	if err != nil {
		return
	}

	return
}
func (srv *SrvCategoriesImpl) SrvFindByid(ctx context.Context, id uint64) (resCategories modelsCategories.Categories, err error) {
	fmt.Println("SrvFindByid")
	resCategories, err = srv.CategoriesRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}

	return
}
func (srv *SrvCategoriesImpl) SrvCreate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error) {
	fmt.Println("SrvCreate")
	tnow := time.Now()
	categoriesIn.Created_At = &tnow
	resCategories, err = srv.CategoriesRepo.RepoCreate(ctx, categoriesIn)
	if err != nil {
		return
	}

	return
}
func (srv *SrvCategoriesImpl) SrvUpdate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error) {
	fmt.Println("SrvUpdate")
	_, err = srv.CategoriesRepo.RepoFindByid(ctx, categoriesIn.Id)
	if err != nil {
		return
	}
	tNow := time.Now()
	categoriesIn.Updated_At = &tNow
	resCategories, err = srv.CategoriesRepo.RepoUpdate(ctx, categoriesIn)
	if err != nil {
		return
	}
	return
}
func (srv *SrvCategoriesImpl) SrvDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("SrvDelete")
	_, err = srv.CategoriesRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	err = srv.CategoriesRepo.RepoDelete(ctx, id)
	if err != nil {
		return
	}

	return
}
