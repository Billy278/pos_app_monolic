package models

import "time"

type Products struct {
	Id          uint64
	Name        string  `json:"name" validate:"required"`
	Stock       uint64  `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Category_id uint64  `json:"category_id" validate:"required"`
	Created_At  *time.Time
	Updated_At  *time.Time
}
