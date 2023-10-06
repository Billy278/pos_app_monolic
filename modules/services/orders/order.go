package services

import (
	"context"

	modelsOrder "github.com/Billy278/pos_app_monolic/modules/models/orders"
)

type OrderSrv interface {
	SrvDetail(ctx context.Context, id uint64) (resOrder modelsOrder.Order, err error)
	SrvList(ctx context.Context, limit, offset uint64) (resOrder []modelsOrder.Order, err error)

	SrvCreate(ctx context.Context, orderIn modelsOrder.Order) (resOrder modelsOrder.Order, err error)
	SrvCreates(ctx context.Context, orderIn modelsOrder.Order) (resOrder modelsOrder.Order, err error)
}
