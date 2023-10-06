package models

import "time"

type OrderProduct struct {
	Id         uint64
	OrderId    uint64
	ProductId  uint64 `json:"product_id" validate:"required,number"`
	Qty        uint64 `json:"qty" validate:"required,number"`
	TotalPrize float64
	Created_At *time.Time
	Updated_At *time.Time
}
