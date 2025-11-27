package main

import "github.com/labstack/echo/v4"

func (a *Application) routes() *echo.Echo {
	e := echo.New()
	a.middleware(e)

	e.Static("/static", "ui/static")

	for _, handler := range a.Handlers {
		handler.RegisterRoutes(e)
	}

	return e
}
