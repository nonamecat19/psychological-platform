package server

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/database"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	database.InitPostgres()

	return &FiberServer{
		App: fiber.New(),
	}
}
