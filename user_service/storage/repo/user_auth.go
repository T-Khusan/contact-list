package repo

import (
	"user_service"
)

// Authorization ...
type UserRepoI interface {
	CreateUser(user user_service.User) (int, error)
}
