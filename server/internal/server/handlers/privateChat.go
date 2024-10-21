package handlers

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/model"
	"server/internal/repository"
	"strconv"
)

func GetAllPrivateChatsHandler(c *fiber.Ctx) error {
	privateChatsRepository := repository.NewPrivateChatRepository()
	privateChats, _ := privateChatsRepository.FindAll()

	return c.JSON(privateChats)
}

func CreatePrivateChatHandler(c *fiber.Ctx) error {
	privateChatsRepository := repository.NewPrivateChatRepository()

	data := new(model.PrivateChat)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	privateChat, _ := privateChatsRepository.Create(*data)

	return c.JSON(privateChat)
}

func UpdatePrivateChatHandler(c *fiber.Ctx) error {
	privateChatsRepository := repository.NewPrivateChatRepository()

	data := new(model.PrivateChat)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	privateChat, _ := privateChatsRepository.UpdateOne(*data)

	return c.JSON(privateChat)
}

func DeletePrivateChatHandler(c *fiber.Ctx) error {
	privateChatsRepository := repository.NewPrivateChatRepository()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = privateChatsRepository.DeleteOne(uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
