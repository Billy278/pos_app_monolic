package respository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	modelsPayment "github.com/Billy278/pos_app_monolic/modules/models/payment"
)

type PaymentImpl struct {
	DB *sql.DB
}

func NewPaymentImpl(db *sql.DB) Payment {
	return &PaymentImpl{
		DB: db,
	}
}

func (repo *PaymentImpl) RepoList(ctx context.Context, limit, offset uint64) (resPayment []modelsPayment.Payment, err error) {
	fmt.Println("RepoList")

	sqlList := "SELECT id,name,type,created_at,updated_at FROM payments Limit $1 offset $2"
	rows, err := repo.DB.QueryContext(ctx, sqlList, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()
	payment := modelsPayment.Payment{}
	for rows.Next() {
		err = rows.Scan(&payment.Id, &payment.Name, &payment.Type, &payment.Created_At, &payment.Updated_At)
		if err != nil {
			return
		}
		resPayment = append(resPayment, payment)
	}
	return
}
func (repo *PaymentImpl) RepoFindByid(ctx context.Context, id uint64) (resPayment modelsPayment.Payment, err error) {
	fmt.Println(" RepoFindByid")
	sqlFind := "SELECT id,name,type,created_at,updated_at FROM payments WHERE id=$1"
	rows, err := repo.DB.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(&resPayment.Id, &resPayment.Name, &resPayment.Type, &resPayment.Created_At, &resPayment.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")
	}
	return
}
func (repo *PaymentImpl) RepoCreate(ctx context.Context, paymentIn modelsPayment.Payment) (resPayment modelsPayment.Payment, err error) {
	fmt.Println("RepoCreate")
	sqlCreate := "INSERT INTO payments(name,type,created_at) VALUES($1,$2,$3) RETURNING id"
	rows, err := repo.DB.QueryContext(ctx, sqlCreate, paymentIn.Name, paymentIn.Type, paymentIn.Created_At)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&resPayment.Id)
		if err != nil {
			return
		}
	}
	resPayment.Name = paymentIn.Name
	resPayment.Type = paymentIn.Type
	resPayment.Created_At = paymentIn.Created_At
	return
}
func (repo *PaymentImpl) RepoUpdate(ctx context.Context, paymentIn modelsPayment.Payment) (resPayment modelsPayment.Payment, err error) {
	fmt.Println("RepoUpdate")
	sqlUpdate := "UPDATE payments set name=$1,type=$2, updated_at=$3 WHERE id=$4"
	_, err = repo.DB.ExecContext(ctx, sqlUpdate, paymentIn.Name, paymentIn.Type, paymentIn.Updated_At, paymentIn.Id)
	if err != nil {
		return
	}
	resPayment.Id = paymentIn.Id
	resPayment.Name = paymentIn.Name
	resPayment.Type = paymentIn.Type
	resPayment.Updated_At = paymentIn.Updated_At
	return
}
func (repo *PaymentImpl) RepoDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("RepoDelete")
	sqlDelete := "DELETE FROM payments WHERE id=$1"
	_, err = repo.DB.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return
	}
	return
}
