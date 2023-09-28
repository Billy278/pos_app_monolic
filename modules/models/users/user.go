package models

import "time"

type User struct {
	Id         uint64
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Created_At *time.Time
	Updated_At *time.Time
}
