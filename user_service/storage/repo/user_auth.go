package repo

import "user_service/genproto/user_service"

// Authorization ...
type UserRepoI interface {
	CreateUser(req *user_service.User) (string, error)
	GetUser(name, password string) (*user_service.User, error)
}
