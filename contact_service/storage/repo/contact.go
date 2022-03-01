package repo

import (
	"contact_service/genproto/contact_service"
	"context"
)

type ContactRepoI interface {
	Create(ctx context.Context, req *contact_service.Contact) (string, error)
	GetAll(req *contact_service.GetAllContactRequest) (*contact_service.GetAllContactResponse, error)
	Get(id string) (*contact_service.Contact, error)
	Update(req *contact_service.Contact) (string, error)
	Delete(id string) (string, error)
}
