package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"user_service/genproto/user_service"
	"user_service/pkg/helper"
	"user_service/pkg/logger"
	"user_service/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
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
func (s *userService) CreateUser(ctx context.Context, req *user_service.User) (*user_service.UserId, error) {
	req.Password = hashPassword(req.Password)
	id, err := s.storage.User().CreateUser(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while create user", req, codes.Internal)
	}

	return &user_service.UserId{
		Id: id,
	}, nil
}

func hashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
