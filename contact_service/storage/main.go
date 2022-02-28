package storage

import (
	"position_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Profession() repo.ProfessionRepoI
	Position() repo.PositionRepoI
	Attribute() repo.AttributeRepoI
}

type storagePg struct {
	db         *sqlx.DB
	profession repo.ProfessionRepoI
	position   repo.PositionRepoI
	attribute  repo.AttributeRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db: db,
	}
}
