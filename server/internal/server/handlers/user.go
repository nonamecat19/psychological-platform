package handlers

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/utils"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()
	users, _ := usersRepository.FindAll()

	return c.JSON(users)
}

func CreateUserHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()

	data := model.User{
		Name:        nil,
		Surname:     nil,
		IsAnonymous: false,
		Role:        "admin",
		Email:       "admin@gmail.com",
		Password:    "admin",
	}

	user, _ := usersRepository.Create(data)

	return c.JSON(user)
}

func UpdateUserHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()
	data := model.User{
		Name: utils.StringPtr("Test"),
	}
	user, err := usersRepository.UpdateOne(1, data)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()

	err := usersRepository.DeleteOne(1)
	if err != nil {
		return err
	}

	return nil
}
