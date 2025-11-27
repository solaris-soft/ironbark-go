package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render replaces Echo's echo.Context.Render() method.
func Render(ctx echo.Context, stausCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(stausCode, buf.String())
}
