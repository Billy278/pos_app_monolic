package sercives

import (
	"context"

	modelsPayment "github.com/Billy278/pos_app_monolic/modules/models/payment"
)

type PaymentSrv interface {
	SrvList(ctx context.Context, limit, offset uint64) (resPayment []modelsPayment.Payment, err error)
	SrvFindByid(ctx context.Context, id uint64) (resPayment modelsPayment.Payment, err error)
	SrvCreate(ctx context.Context, paymentIn modelsPayment.Payment) (resPayment modelsPayment.Payment, err error)
	SrvUpdate(ctx context.Context, paymentIn modelsPayment.Payment) (resPayment modelsPayment.Payment, err error)
	SrvDelete(ctx context.Context, id uint64) (err error)
}
