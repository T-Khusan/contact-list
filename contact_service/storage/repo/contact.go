package repo

import (
	"contact_service/genproto/contact_service"
	"context"
)

type ContactRepoI interface {
	Create(ctx context.Context, req *contact_service.Contact) (string, error)
	GetAll(req *contact_service.UserId) (*contact_service.GetAllContactResponse, error)
	Get(req *contact_service.ContactUserId) (*contact_service.Contact, error)
	Update(req *contact_service.Contact) (*contact_service.ContactUpdate, error)
	Delete(req *contact_service.ContactUserId) (string, error)
}
