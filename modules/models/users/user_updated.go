package models

type UserUpdated struct {
	Id    uint64
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
