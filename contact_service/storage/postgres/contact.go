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

func (r *contactRepo) GetAll(req *contact_service.GetAllContactRequest) (*contact_service.GetAllContactResponse, error) {
	var (
		filter      string
		args        = make(map[string]interface{})
		count       int32
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
		Count:       count,
	}, nil

}
