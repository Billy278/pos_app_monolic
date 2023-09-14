package respository

import (
	"context"

	modelsPayment "github.com/Billy278/pos_app_monolic/modules/models/payment"
)

type Payment interface {
	RepoList(ctx context.Context, limit, offset uint64) (resPayment []modelsPayment.Payment, err error)
	RepoFindByid(ctx context.Context, id uint64) (resPayment modelsPayment.Payment, err error)
	RepoCreate(ctx context.Context, paymentIn modelsPayment.Payment) (resPayment modelsPayment.Payment, err error)
	RepoUpdate(ctx context.Context, paymentIn modelsPayment.Payment) (resPayment modelsPayment.Payment, err error)
	RepoDelete(ctx context.Context, id uint64) (err error)
}
