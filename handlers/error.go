package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/solaris-soft/ironbark-go/ui/pages"
)

// Error renders an error page.
func Error(ctx echo.Context, statusCode int, message string) error {
	return Render(ctx, statusCode, pages.Error(statusCode, message))
}
