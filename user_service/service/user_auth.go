package service

import (
	"crypto/sha1"
	"fmt"
	"user_service"
	"user_service/pkg/logger"
	"user_service/storage"

	"github.com/jmoiron/sqlx"
)

type userService struct {
	logger  logger.Logger
	storage storage.StorageI
}

const (
	salt = "aSWdkh6465a4dEWdyKHJS"
)


func NewUserService(db *sqlx.DB, log logger.Logger) *userService {
	return &userService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

// CreateUser ...
func (s *userService) CreateUser(user user_service.User) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.storage.User().CreateUser(user)
}

func hashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
