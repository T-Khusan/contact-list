package storage

import (
	"user_service/storage/postgres"
	"user_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserRepoI
}

type storagePg struct {
	db   *sqlx.DB
	user repo.UserRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:      db,
		user: postgres.NewContactRepo(db),
	}
}

func (s *storagePg) User() repo.UserRepoI {
	return s.user // check nil
}
