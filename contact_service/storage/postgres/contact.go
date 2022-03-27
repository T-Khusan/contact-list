package postgres

import (
	"contact_service/genproto/contact_service"
	"contact_service/storage/repo"
	"context"
	"fmt"
	"strings"

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
				phone,
				user_id
			) 
			VALUES ($1, $2, $3) RETURNING id`

	if err := tx.QueryRowContext(ctx, query, req.Name, req.Phone, req.UserId).Scan(&id); err != nil {
		if err := tx.Rollback(); err != nil {
			return "", err
		}
	}

	return id, tx.Commit()
}

func (r *contactRepo) GetAll(req *contact_service.UserId) (*contact_service.Contact, error) {
	var contacts []*contact_service.Contact

	query := "SELECT id, name, phone, user_id FROM contact WHERE user_id=$1"

	err := r.db.Select(&contacts, query, user_id)

	return &contact_service.Contact{
		Id:     req.Id,
		Name:   req.Name,
		Phone:  req.Phone,
		UserId: req.UserId,
	}, nil

}

func (r *contactRepo) Get(userID, contactID string) (*contact_service.Contact, error) {
	var contact contact_service.Contact

	query := `SELECT id, name, phone, user_id FROM contact WHERE user_id = $1 AND id=$2`

	row := r.db.QueryRow(query, userID, contactID)
	err := row.Scan(
		&contact.UserId,
		&contact.Name,
		&contact.Phone,
		&contact.Id,
	)

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *contactRepo) Update(req *contact_service.Contact) (string, error) {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if req.Name != "" {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argID))
		args = append(args, req.Name)
		argID++
	}

	if req.Phone != "" {
		setValue = append(setValue, fmt.Sprintf("phone=$%d", argID))
		args = append(args, req.Phone)
		argID++
	}

	argString := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE contact SET %s WHERE user_id=$%d AND id=$%d", argString, argID, argID+1)

	args = append(args, req.UserId, req.Id)

	_, err := r.db.Exec(query, args...)

	return "Updated successfuly", err
}

func (r *contactRepo) Delete(userID, contactID string) (string, error) {
	query := "DELETE FROM contact WHERE user_id=$1 AND id=$2"

	_, err := r.db.Exec(query, query, userID, contactID)


	if err != nil {
		return "", err
	}

	return "Deleted one row successfuly", nil
}
