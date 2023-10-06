package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	models "github.com/Billy278/pos_app_monolic/modules/models/order_product"
	modelsOrder "github.com/Billy278/pos_app_monolic/modules/models/orders"
	modelsPayment "github.com/Billy278/pos_app_monolic/modules/models/payment"
)

type OrderRepoImpl struct {
}

func NewOrderRepoImpl() OrderRepo {
	return &OrderRepoImpl{}

}
func (repo *OrderRepoImpl) RepoDetail(ctx context.Context, db *sql.DB, id uint64) (resOrder modelsOrder.Order, err error) {
	fmt.Println("RepoDetail")
	sqlList := `SELECT o.id,o.user_id,o.payment_id,o.total_price,o.total_paid,o.total_return, o.created_at,o.updated_at,
	p.id,p.order_id,p.product_id,p.qty,p.total_price,p.created_at,p.updated_at
	 FROM orders AS o
	 INNER JOIN order_product AS p
	 ON o.id=p.order_id
	 WHERE o.id=$1
	`
	row, err := db.QueryContext(ctx, sqlList, id)
	if err != nil {
		return
	}
	defer row.Close()
	product := models.OrderProduct{}
	tempOrderProduct := []models.OrderProduct{}

	for row.Next() {

		err = row.Scan(&resOrder.Id, &resOrder.UserId, &resOrder.PaymentId, &resOrder.TotalPrize, &resOrder.TotalPaid, &resOrder.TotalReturn, &resOrder.Created_At, &resOrder.Updated_At, &product.Id, &product.OrderId, &product.ProductId, &product.Qty, &product.TotalPrize, &product.Created_At, &product.Updated_At)
		if err != nil {
			return
		}
		tempOrderProduct = append(tempOrderProduct, product)

	}
	resOrder.ProductDetail = &tempOrderProduct
	return
}

func (repo *OrderRepoImpl) RepoList(ctx context.Context, db *sql.DB, limit, offset uint64) (resOrder []modelsOrder.Order, err error) {
	fmt.Println("RepoOrderList")
	sqlList := `SELECT o.id,o.user_id,o.payment_id,o.total_price,o.total_paid,o.total_return, o.created_at,o.updated_at,
	py.id, py.name,py.type,py.created_at,py.updated_at
	FROM orders AS o
	INNER JOIN payments AS py
	ON o.payment_id=py.id
	 LIMIT $1 OFFSET $2`
	row, err := db.QueryContext(ctx, sqlList, limit, offset)
	if err != nil {
		return
	}
	defer row.Close()
	order := modelsOrder.Order{}
	payment := modelsPayment.Payment{}

	for row.Next() {
		err = row.Scan(&order.Id, &order.UserId, &order.PaymentId, &order.TotalPrize, &order.TotalPaid, &order.TotalReturn, &order.Created_At, &order.Updated_At, &payment.Id, &payment.Name, &payment.Type, &payment.Created_At, &payment.Updated_At)
		if err != nil {
			return
		}
		order.PaymentDetail = &payment
		resOrder = append(resOrder, order)
	}

	return
}

// fungsi unusing
func (repo *OrderRepoImpl) RepoListUn(ctx context.Context, db *sql.DB, limit, offset uint64) (resOrder []modelsOrder.Order, err error) {
	fmt.Println("RepoOrderList")
	sqlList := `SELECT o.id,o.user_id,o.payment_id,o.total_price,o.total_paid,o.total_return, o.created_at,o.updated_at,
	p.id,p.order_id,p.product_id,p.qty,p.total_price,p.created_at,p.updated_at
	 FROM orders AS o
	 INNER JOIN order_product AS p
	 ON o.id=p.order_id`
	row, err := db.QueryContext(ctx, sqlList)
	if err != nil {
		return
	}
	defer row.Close()
	order := modelsOrder.Order{}
	product := models.OrderProduct{}
	tempOrderProduct := []models.OrderProduct{}
	for row.Next() {
		err = row.Scan(&order.Id, &order.UserId, &order.PaymentId, &order.TotalPrize, &order.TotalPaid, &order.TotalReturn, &order.Created_At, &order.Updated_At, &product.Id, &product.OrderId, &product.ProductId, &product.Qty, &product.TotalPrize, &product.Created_At, &product.Updated_At)
		if err != nil {
			return
		}
		tempOrderProduct = append(tempOrderProduct, product)
		order.ProductDetail = &tempOrderProduct
		resOrder = append(resOrder, order)
	}

	return
}
func (repo *OrderRepoImpl) RepoFindByid(ctx context.Context, db *sql.DB, id uint64) (resOrder modelsOrder.Order, err error) {
	fmt.Println("RepoOrderFindByid")
	sqlFind := "SELECT id,user_id,payment_id,total_price,total_paid,total_return,created_at,updated_at FROM orders WHERE id=$1"
	row, err := db.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resOrder.Id, &resOrder.UserId, &resOrder.PaymentId, &resOrder.TotalPrize, &resOrder.TotalPaid, &resOrder.TotalReturn, &resOrder.Created_At, &resOrder.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("ORDER NOT FOUND")
	}
	return

}
func (repo *OrderRepoImpl) RepoCreate(ctx context.Context, tx *sql.Tx, orderIn modelsOrder.Order) (resOrder modelsOrder.Order, err error) {
	fmt.Println("RepoOrderCreate")
	sqlCreate := "INSERT INTO orders(user_id,payment_id,total_price,total_paid,total_return, created_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id"
	row, err := tx.QueryContext(ctx, sqlCreate, orderIn.UserId, orderIn.PaymentId, orderIn.TotalPrize, orderIn.TotalPaid, orderIn.TotalReturn, orderIn.Created_At)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resOrder.Id)
		if err != nil {
			return
		}
	}

	resOrder.UserId = orderIn.UserId
	resOrder.PaymentId = orderIn.PaymentId
	resOrder.TotalPrize = orderIn.TotalPrize
	resOrder.TotalPaid = orderIn.TotalPaid
	resOrder.TotalReturn = orderIn.TotalReturn
	resOrder.Created_At = orderIn.Created_At

	return
}
func (repo *OrderRepoImpl) RepoDelete(ctx context.Context, db *sql.DB, id uint64) (err error) {
	fmt.Println("RepoOrderDelete")
	sqlDelete := "DELETE FROM orders WHERE id=$1"
	_, err = db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return
	}
	return
}
