package models

// User struct
type UserModel struct {
	ID       string `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SigninInput struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
