package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	modelsProducts "github.com/Billy278/pos_app_monolic/modules/models/products"
	repository "github.com/Billy278/pos_app_monolic/modules/repository/products"
	services "github.com/Billy278/pos_app_monolic/modules/services/categories"
)

type SrvProductImpl struct {
	ProductRepo repository.Product
	CategorySrv services.SrvCategories
}

func NewSrvProductImpl(productrepo repository.Product, categorysrv services.SrvCategories) SrvProduct {
	return &SrvProductImpl{
		ProductRepo: productrepo,
		CategorySrv: categorysrv,
	}
}

func (srv *SrvProductImpl) SrvList(ctx context.Context, limit, offset uint64) (resProduct []modelsProducts.Products, err error) {
	fmt.Println("SrvList")
	resProduct, err = srv.ProductRepo.RepoList(ctx, limit, offset)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvFindByid(ctx context.Context, id uint64) (resProduct modelsProducts.Products, err error) {
	fmt.Println("SrvFindByid")
	resProduct, err = srv.ProductRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvCreate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error) {
	fmt.Println("SrvCreate")
	_, err = srv.CategorySrv.SrvFindByid(ctx, productIn.Category_id)
	if err != nil {
		err = errors.New("CATEGORY ID IS NOT FOUND")
		return
	}
	tNow := time.Now()
	productIn.Created_At = &tNow
	resProduct, err = srv.ProductRepo.RepoCreate(ctx, productIn)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvUpdate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error) {
	fmt.Println("SrvUpdate")
	_, err = srv.CategorySrv.SrvFindByid(ctx, productIn.Category_id)
	if err != nil {
		err = errors.New("CATEGORY ID IS NOT FOUND")
		return
	}
	_, err = srv.ProductRepo.RepoFindByid(ctx, productIn.Id)
	if err != nil {
		return
	}
	tNow := time.Now()
	productIn.Updated_At = &tNow
	resProduct, err = srv.ProductRepo.RepoUpdate(ctx, productIn)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("SrvDelete")
	_, err = srv.ProductRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	err = srv.ProductRepo.RepoDelete(ctx, id)
	if err != nil {
		return
	}
	return
}
