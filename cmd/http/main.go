package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/solaris-soft/ironbark-go/config"
	"github.com/solaris-soft/ironbark-go/db"
	"github.com/solaris-soft/ironbark-go/handlers"
	"github.com/solaris-soft/ironbark-go/services"
)

// Handler is an interface for handlers that can register routes.
type Handler interface {
	RegisterRoutes(e *echo.Echo)
}

// Application holds the server state
type Application struct {
	DB       *db.Queries
	Handlers []Handler
}

// Start the application.
func (a Application) start(port string) {
	e := a.routes()
	e.Logger.Fatal(e.Start(port))
}

func main() {
	config := config.LoadConfig()
	database, err := pgx.Connect(context.Background(), config.DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close(context.Background())
	queries := db.New(database)

	handlers := []Handler{
		handlers.NewContactsHandler(queries, services.NewContactService(queries)),
		handlers.NewIndexHandler(),
	}

	a := Application{
		queries,
		handlers,
	}
	a.start(":8080")
}
