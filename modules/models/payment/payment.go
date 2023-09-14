package models

import "time"

type Payment struct {
	Id         uint64
	Name       string `json:"name" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Created_At *time.Time
	Updated_At *time.Time
}
