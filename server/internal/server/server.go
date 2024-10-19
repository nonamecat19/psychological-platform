package server

import (
	"github.com/gofiber/fiber/v2"

	"server/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(),

		db: database.New(),
	}

	return server
}