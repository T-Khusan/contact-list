package service

import (
	"contact_service/genproto/contact_service"
	"contact_service/pkg/helper"
	"contact_service/pkg/logger"
	"contact_service/storage"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
)

type contactService struct {
	logger  logger.Logger
	storage storage.StorageI
	contact_service.UnimplementedContactServiceServer
}

func NewContactService(db *sqlx.DB, log logger.Logger) *contactService {
	return &contactService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *contactService) Create(ctx context.Context, req *contact_service.Contact) (*contact_service.ContactId, error) {
	fmt.Println(s.storage)
	id, err := s.storage.Contact().Create(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while create contact", req, codes.Internal)
	}

	return &contact_service.ContactId{
		Id: id,
	}, nil
}

func (s *contactService) GetAll(ctx context.Context, req *contact_service.GetAllContactRequest) (*contact_service.GetAllContactResponse, error) {
	resp, err := s.storage.Contact().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all contacts", req, codes.Internal)
	}

	return resp, nil
}
