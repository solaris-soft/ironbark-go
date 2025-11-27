package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solaris-soft/ironbark-go/db"
	"github.com/solaris-soft/ironbark-go/serializers"
	pages "github.com/solaris-soft/ironbark-go/ui/pages/contacts"
)

// ContactService is a service for the contacts resource.
type ContactService interface {
	ListContacts(ctx context.Context) ([]db.Contact, error)
}

// ContactsHandler is a handler for the contacts resource.
type ContactsHandler struct {
	DB      *db.Queries
	Service ContactService
}

// NewContactsHandler creates a new ContactsHandler.
func NewContactsHandler(db *db.Queries, service ContactService) *ContactsHandler {
	return &ContactsHandler{
		DB:      db,
		Service: service,
	}
}

// RegisterRoutes registers the routes for the contacts resource.
func (h *ContactsHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/contacts", h.ListContactsAPI)
	e.GET("/contacts", h.ListContactsPage)
}

// ListContacts lists all contacts.
func (h *ContactsHandler) ListContactsAPI(c echo.Context) error {
	contacts, err := h.Service.ListContacts(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to list contacts"})
	}
	return c.JSON(http.StatusOK, contacts)
}

// ListContactsPage lists all contacts in a page.
func (h *ContactsHandler) ListContactsPage(c echo.Context) error {
	contacts, err := h.Service.ListContacts(c.Request().Context())
	if err != nil {
		return Error(c, http.StatusInternalServerError, "Failed to list contacts")
	}

	contactSerializers := make([]serializers.ContactSerializer, len(contacts))
	for i, contact := range contacts {
		contactSerializers[i] = *serializers.NewContact(&contact)
	}

	return Render(c, http.StatusOK, pages.Index(contactSerializers))
}
