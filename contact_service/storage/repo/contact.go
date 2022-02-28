package repo

import "contact_service/genproto/contact_service"


type ContactRepoI interface {
	Create(req *contact_service.Contact) (string, error)
}
