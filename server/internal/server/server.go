package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"server/internal/database"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	database.InitPostgres()

	app := fiber.New()
	app.Use(cors.New())

	return &FiberServer{
		App: app,
	}
}
