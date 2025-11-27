package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solaris-soft/ironbark-go/ui/layouts"
	pages "github.com/solaris-soft/ironbark-go/ui/pages"
)

type IndexHandler struct{}

// NewIndexHandler creates a new IndexHandler.
func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

// RegisterRoutes registers the routes for the index resource.
func (h *IndexHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.Index)
}

// Index renders the index page.
func (h *IndexHandler) Index(c echo.Context) error {
	return Render(c, http.StatusOK, layouts.BaseLayout("Home", pages.Index()))
}
