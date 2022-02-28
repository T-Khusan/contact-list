package repo

import "contact_service/genproto/contact_service"

// "bitbucket.org/Udevs/position_service/genproto/position_service"

type ContactRepoI interface {
	Create(req *contact_service.Contact) (string, error)
}
