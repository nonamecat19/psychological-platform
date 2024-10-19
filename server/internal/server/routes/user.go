package routes

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server/internal/repository"
	"server/internal/server"
)

func ConnectUsersRouter(s *server.FiberServer) {
	usersRouter := s.Group("/users")

	usersRouter.Get("/", func(c *fiber.Ctx) error {
		var usersRepository = repository.NewUserRepository()
		users, _ := usersRepository.FindAll()

		log.Println(users)

		return c.JSON(fiber.Map{
			"message": users,
		})
	})

	usersRouter.Post("/", func(c *fiber.Ctx) error {
		var usersRepository = repository.NewUserRepository()
		users, _ := usersRepository.FindAll()

		log.Println(users)

		return c.JSON(fiber.Map{
			"create": users,
		})
	})
}
