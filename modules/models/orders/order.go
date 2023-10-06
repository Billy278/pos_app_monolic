package models

import (
	"time"

	modelsProduct "github.com/Billy278/pos_app_monolic/modules/models/order_product"
	modelsPayment "github.com/Billy278/pos_app_monolic/modules/models/payment"
)

type Order struct {
	Id            uint64
	UserId        uint64
	PaymentId     uint64 `json:"payment_id" validate:"required,number"`
	TotalPrize    float64
	TotalPaid     float64 `json:"total_paid" validate:"required,number"`
	TotalReturn   float64
	ProductDetail *[]modelsProduct.OrderProduct `json:"products" validate:"required,dive,required"`
	PaymentDetail *modelsPayment.Payment        `json:"paymentDetail"`
	Created_At    *time.Time
	Updated_At    *time.Time
}
