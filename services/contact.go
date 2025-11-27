package services

import (
	"context"

	"github.com/solaris-soft/ironbark-go/db"
)

type ContactService struct {
	DB *db.Queries
}

// NewContactService creates a new ContactService.
func NewContactService(db *db.Queries) *ContactService {
	return &ContactService{DB: db}
}

// ListContacts lists all contacts.
func (s *ContactService) ListContacts(ctx context.Context) ([]db.Contact, error) {
	return s.DB.ListContacts(ctx)
}
