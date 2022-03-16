package models

type CreateContactModel struct {
	Name string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
