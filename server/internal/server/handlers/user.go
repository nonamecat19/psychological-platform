package handlers

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/model"
	"server/internal/repository"
	"strconv"
)

func GetMeHandler(c *fiber.Ctx) error {
	info := GetUserAuthInfo(c)
	usersRepository := repository.NewUserRepository()
	user, _ := usersRepository.GetOne(info.ID)
	return c.JSON(user)
}

func GetAllUsersHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()
	users, _ := usersRepository.FindAll()

	return c.JSON(users)
}

func GetAllPsychologistsHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()
	users, _ := usersRepository.FindAllPsychologists()

	return c.JSON(users)
}

func CreateUserHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()

	data := new(model.User)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	user, _ := usersRepository.Create(*data)

	return c.JSON(user)
}

func UpdateUserHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()

	data := new(model.User)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	user, _ := usersRepository.UpdateOne(*data)

	return c.JSON(user)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	usersRepository := repository.NewUserRepository()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = usersRepository.DeleteOne(uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
