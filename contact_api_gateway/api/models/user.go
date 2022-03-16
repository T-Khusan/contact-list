package models

// User struct
type UserModel struct {
	ID       string `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
