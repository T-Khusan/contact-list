package postgres

import (
	"contact_service/genproto/contact_service"
	"contact_service/storage/repo"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type contactRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewContactRepo(db *sqlx.DB) repo.ContactRepoI {
	return &contactRepo{db: db}
}

func (r *contactRepo) Create(req *contact_service.Contact) (string, error) {
	var (
		err error
		tx  *sql.Tx
		id  uuid.UUID
	)
	tx, err = r.db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err != nil {
		return "", err
	}

	id, err = uuid.NewRandom()
	if err != nil {
		return "", err
	}

	query := `INSERT INTO contact (
				id,
				name,
				phone
			) 
			VALUES ($1, $2, $3) `

	_, err = tx.Exec(query, id, req.Name, req.Phone)

	if err != nil {
		return "", err
	}
	return id.String(), nil
}
