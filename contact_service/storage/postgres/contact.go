package postgres

import (
	"contact_service/genproto/contact_service"
	"contact_service/storage/repo"
	"context"
	"database/sql"

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

func (r *contactRepo) GetAll(req *contact_service.GetAllContactRequest) (*contact_service.GetAllContactResponse, error) {
	var (
		filter   string
		args     = make(map[string]interface{})
		count    int32
		contacts []*contact_service.Contact
	)

	if req.Name != "" {
		filter += " AND name ilike '%' || :name || '%' "
		args["name"] = req.Name
	}

	countQuery := `SELECT count(1) FROM contact WHERE true ` + filter
	rows, err := r.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	query := `SELECT
					id,
					name,
					phone
				FROM contact WHERE true ` + filter

	rows, err = r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var contact contact_service.Contact

		err = rows.Scan(
			&contact.Id,
			&contact.Name,
			&contact.Phone,
		)

		if err != nil {
			return nil, err
		}

		contacts = append(contacts, &contact)
	}

	return &contact_service.GetAllContactResponse{
		Contacts: contacts,
		Count:    count,
	}, nil

}

func (r *contactRepo) Get(id string) (*contact_service.Contact, error) {
	var contact contact_service.Contact

	query := `SELECT id, name, phone FROM contact WHERE id = $1`

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&contact.Id,
		&contact.Name,
		&contact.Phone,
	)

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *contactRepo) Update(req *contact_service.Contact) (string, error) {
	var (
		err error
		tx  *sql.Tx
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

	query := `
		UPDATE contact
		SET name = $1, phone = $2
		WHERE id = $3
	`

	_, err = tx.Exec(query, req.Name, req.Phone, req.Id)
	if err != nil {
		return "", err
	}

	return "Updated successfuly", nil
}

func (r *contactRepo) Delete(id string) (string, error) {
	query := `DELETE FROM contact WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return "", err
	}

	return "Deleted one row successfuly", nil
}
