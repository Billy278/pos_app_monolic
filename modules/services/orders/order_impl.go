package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	modelsOrder "github.com/Billy278/pos_app_monolic/modules/models/orders"
	modelsProduct "github.com/Billy278/pos_app_monolic/modules/models/products"
	repoDetail "github.com/Billy278/pos_app_monolic/modules/repository/order_product"
	repository "github.com/Billy278/pos_app_monolic/modules/repository/orders"
	repoProduct "github.com/Billy278/pos_app_monolic/modules/repository/products"
)

type OrderSrvImpl struct {
	DB            *sql.DB
	OrderRepo     repository.OrderRepo
	ProductDetail repoDetail.OrderProduct
	ProductRepo   repoProduct.Product
}

func NewOrderSrvImpl(db *sql.DB, orderrepo repository.OrderRepo, productdtl repoDetail.OrderProduct, productrepo repoProduct.Product) OrderSrv {
	return &OrderSrvImpl{
		DB:            db,
		OrderRepo:     orderrepo,
		ProductDetail: productdtl,
		ProductRepo:   productrepo,
	}
}

func (srv *OrderSrvImpl) SrvDetail(ctx context.Context, id uint64) (resOrder modelsOrder.Order, err error) {
	fmt.Println("SrvDetail")
	resOrder, err = srv.OrderRepo.RepoDetail(ctx, srv.DB, id)
	if err != nil {
		return
	}
	return
}
func (srv *OrderSrvImpl) SrvList(ctx context.Context, limit, offset uint64) (resOrder []modelsOrder.Order, err error) {
	fmt.Println("SrvList")
	resOrder, err = srv.OrderRepo.RepoList(ctx, srv.DB, limit, offset)
	if err != nil {
		return
	}
	return
}

func (srv *OrderSrvImpl) SrvCreates(ctx context.Context, orderIn modelsOrder.Order) (resOrder modelsOrder.Order, err error) {
	fmt.Println("SrvOrderCreate")
	tx, err := srv.DB.Begin()
	if err != nil {
		return
	}
	// validasi id product dan pastikan stock >=jumlah order
	lenProduct := len(*orderIn.ProductDetail)
	tempProduct := *orderIn.ProductDetail
	var result float64
	var tempResProduct []modelsProduct.Products

	for i := 0; i < lenProduct-1; i++ {
		tempResProduct[i], err = srv.ProductRepo.RepoFindByidTx(ctx, tx, tempProduct[i].Id)
		if err != nil {
			tx.Rollback()
			return
		}
		if tempResProduct[i].Stock < tempProduct[i].Qty {
			msg := fmt.Sprintf("Stock Product with id=%v tidak Mencukupi", tempResProduct[i].Id)
			err = errors.New(msg)
			tx.Rollback()
			return
		}
		tempResProduct[i].Stock = tempResProduct[i].Stock - tempProduct[i].Qty
		tempProduct[i].TotalPrize = tempResProduct[i].Price * float64(tempProduct[i].Qty)
		result = result + (tempResProduct[i].Price * float64(tempProduct[i].Qty))
	}
	// cek request paid ? >= result
	if orderIn.TotalPaid < result {
		tx.Rollback()
		return
	}
	//updated data product
	tNow := time.Now()

	for i := 0; i < lenProduct-1; i++ {
		tempResProduct[i].Updated_At = &tNow
		_, err = srv.ProductRepo.RepoUpdateTx(ctx, tx, tempResProduct[i])
		if err != nil {
			tx.Rollback()
			return
		}
	}

	// create order product
	for i := 0; i < lenProduct-1; i++ {
		tempProduct[i].Created_At = &tNow
		tempProduct[i], err = srv.ProductDetail.RepoCreate(ctx, tx, tempProduct[i])
		if err != nil {
			tx.Rollback()
			return
		}
	}
	//create data order
	orderIn.Created_At = &tNow
	resOrder, err = srv.OrderRepo.RepoCreate(ctx, tx, orderIn)
	if err != nil {
		tx.Rollback()
		return
	}
	resOrder.ProductDetail = &tempProduct
	tx.Commit()
	return
}

func (srv *OrderSrvImpl) SrvCreate(ctx context.Context, orderIn modelsOrder.Order) (resOrder modelsOrder.Order, err error) {
	fmt.Println("SrvOrderCreate")
	tx, err := srv.DB.Begin()
	if err != nil {
		return
	}
	// validasi id product dan pastikan stock >=jumlah order
	lenProduct := len(*orderIn.ProductDetail)
	tempProduct := *orderIn.ProductDetail
	var result float64
	var tempResProduct []modelsProduct.Products
	var temp modelsProduct.Products
	for i := 0; i < lenProduct; i++ {
		temp, err = srv.ProductRepo.RepoFindByidTx(ctx, tx, tempProduct[i].ProductId)
		if err != nil {
			tx.Rollback()
			return
		}
		tempResProduct = append(tempResProduct, temp)

		if tempResProduct[i].Stock < tempProduct[i].Qty {
			msg := fmt.Sprintf("Stock Product with id=%v tidak Mencukupi", tempResProduct[i].Id)
			err = errors.New(msg)
			tx.Rollback()
			return
		}

		tempResProduct[i].Stock = tempResProduct[i].Stock - tempProduct[i].Qty
		tempProduct[i].TotalPrize = tempResProduct[i].Price * float64(tempProduct[i].Qty)
		result = result + (tempResProduct[i].Price * float64(tempProduct[i].Qty))
	}

	// cek request paid ? >= result
	if orderIn.TotalPaid < result {
		tx.Rollback()
		return
	}
	//updated data product
	tNow := time.Now()

	for i := 0; i < lenProduct; i++ {
		tempResProduct[i].Updated_At = &tNow
		_, err = srv.ProductRepo.RepoUpdateTx(ctx, tx, tempResProduct[i])
		if err != nil {
			tx.Rollback()
			return
		}
	}

	//create data order
	orderIn.Created_At = &tNow
	orderIn.TotalPrize = result
	orderIn.TotalReturn = orderIn.TotalPaid - result
	resOrder, err = srv.OrderRepo.RepoCreate(ctx, tx, orderIn)
	if err != nil {
		tx.Rollback()
		return
	}
	// create order product
	for i := 0; i < lenProduct; i++ {
		tempProduct[i].Created_At = &tNow
		tempProduct[i].OrderId = resOrder.Id
		tempProduct[i], err = srv.ProductDetail.RepoCreate(ctx, tx, tempProduct[i])
		if err != nil {
			tx.Rollback()
			return
		}
	}
	resOrder.ProductDetail = &tempProduct
	tx.Commit()
	return
}
