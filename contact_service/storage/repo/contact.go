package repo

import (
	"contact_service/genproto/contact_service"
	"context"
)

type ContactRepoI interface {
	Create(ctx context.Context, req *contact_service.Contact) (string, error)
	GetAll(req *contact_service.UserId) (*contact_service.Contact, error)
	Get(req *contact_service.ContactId) (*contact_service.Contact, error)
	Update(req *contact_service.Contact) (string, error)
	Delete(req *contact_service.ContactId) (string, error)
}
