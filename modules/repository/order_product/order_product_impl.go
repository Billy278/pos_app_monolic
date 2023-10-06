package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/Billy278/pos_app_monolic/modules/models/order_product"
)

type OrderproductImpl struct {
}

func NewOrderproductImpl() OrderProduct {
	return &OrderproductImpl{}
}

func (repo *OrderproductImpl) RepoList(ctx context.Context, db *sql.DB, id uint64) (resProduct []models.OrderProduct, err error) {
	fmt.Println("Repo OrderProductDetail")
	sqlList := "SELECT id,order_id,product_id,qty,total_price,created_at,updated_at FROM order_product WHERE order_id=$1"
	row, err := db.QueryContext(ctx, sqlList, id)
	if err != nil {
		return
	}
	defer row.Close()
	order := models.OrderProduct{}
	for row.Next() {
		err = row.Scan(&order.Id, &order.OrderId, &order.ProductId, &order.Qty, &order.TotalPrize, &order.Created_At, &order.Updated_At)
		if err != nil {
			return
		}
		resProduct = append(resProduct, order)
	}
	return
}
func (repo *OrderproductImpl) RepoCreate(ctx context.Context, tx *sql.Tx, productIn models.OrderProduct) (resProduct models.OrderProduct, err error) {
	fmt.Println("RepoOrderProductDetailCreate")
	fmt.Println(productIn)
	sqlCreate := "INSERT INTO order_product(order_id,product_id,qty,total_price,created_at) VALUES($1,$2,$3,$4,$5) RETURNING id"
	row, err := tx.QueryContext(ctx, sqlCreate, productIn.OrderId, productIn.ProductId, productIn.Qty, productIn.TotalPrize, productIn.Created_At)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resProduct.Id)
		if err != nil {
			return
		}
	}
	resProduct.OrderId = productIn.OrderId
	resProduct.ProductId = productIn.ProductId
	resProduct.Qty = productIn.Qty
	resProduct.TotalPrize = productIn.TotalPrize
	resProduct.Created_At = productIn.Created_At

	return

}
