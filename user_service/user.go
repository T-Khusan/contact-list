package user_service

// User struct
type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
