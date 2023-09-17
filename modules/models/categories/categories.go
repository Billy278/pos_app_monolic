package models

import "time"

type Categories struct {
	Id         uint64
	Name       string `json:"name" validate:"required"`
	Created_At *time.Time
	Updated_At *time.Time
}
