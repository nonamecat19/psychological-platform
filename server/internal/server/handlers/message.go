package handlers

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/model"
	"server/internal/repository"
	"strconv"
)

func GetAllMessagesHandler(c *fiber.Ctx) error {
	messagesRepository := repository.NewMessageRepository()
	messages, _ := messagesRepository.FindAll()

	return c.JSON(messages)
}

func CreateMessageHandler(c *fiber.Ctx) error {
	messagesRepository := repository.NewMessageRepository()

	data := new(model.Message)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	message, _ := messagesRepository.Create(*data)

	return c.JSON(message)
}

func UpdateMessageHandler(c *fiber.Ctx) error {
	messagesRepository := repository.NewMessageRepository()

	data := new(model.Message)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	message, _ := messagesRepository.UpdateOne(*data)

	return c.JSON(message)
}

func DeleteMessageHandler(c *fiber.Ctx) error {
	messagesRepository := repository.NewMessageRepository()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = messagesRepository.DeleteOne(uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
