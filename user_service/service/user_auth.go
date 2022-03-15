package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"user_service/genproto/user_service"
	"user_service/pkg/helper"
	"user_service/pkg/logger"
	"user_service/storage"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
)

type userService struct {
	logger  logger.Logger
	storage storage.StorageI
}

const (
	salt       = "aSWdkh6465a4dEWdyKHJS"
	timeToken  = 12 * time.Hour
	signingKey = "asd12aHJGJHG4sad"
)

type tokenClaims struct {
	UserID string
	jwt.StandardClaims
}

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

// GenerateToken ...
func (s *userService) GenerateToken(username, password string) (string, error) {
	user, err := s.storage.User().GetUser(username, hashPassword(password))
	if err != nil {
		return "", err
	}

	tk := tokenClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeToken).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	return token.SignedString([]byte(signingKey))
}

// ParseToken token parse and returen user id
func (s *userService) ParseToken(token string) (string, error) {
	tk, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return "", nil
	}

	cl, ok := tk.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return cl.UserID, nil
}

func hashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
