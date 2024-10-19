package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server/internal/repository"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	var usersRepository = repository.NewUserRepository()
	users, _ := usersRepository.FindAll()

	log.Println(users)

	return c.JSON(fiber.Map{
		"message": users,
	})
}
