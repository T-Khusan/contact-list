package repo

import (
	"contact_service/genproto/contact_service"
	"context"
)

type ContactRepoI interface {
	Create(ctx context.Context, req *contact_service.Contact) (string, error)
}
