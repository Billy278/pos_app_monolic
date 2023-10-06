package repository

import (
	"context"
	"database/sql"

	models "github.com/Billy278/pos_app_monolic/modules/models/order_product"
)

type OrderProduct interface {
	RepoList(ctx context.Context, db *sql.DB, id uint64) (resProduct []models.OrderProduct, err error)
	RepoCreate(ctx context.Context, tx *sql.Tx, productIn models.OrderProduct) (resProduct models.OrderProduct, err error)
}
