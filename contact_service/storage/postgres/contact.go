package postgres

import (
	"contact_service/genproto/contact_service"
	"contact_service/storage/repo"
	"context"

	"github.com/jmoiron/sqlx"
)

type contactRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewContactRepo(db *sqlx.DB) repo.ContactRepoI {
	return &contactRepo{db: db}
}

func (r *contactRepo) Create(ctx context.Context, req *contact_service.Contact) (string, error) {
	var id string

	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	query := `INSERT INTO contact (
				name,
				phone
			) 
			VALUES ($1, $2) RETURNING id`

	if err := tx.QueryRowContext(ctx, query, req.Name, req.Phone).Scan(&id); err != nil {
		if err := tx.Rollback(); err != nil {
			return "", err
		}
	}

	return id, tx.Commit()
}
