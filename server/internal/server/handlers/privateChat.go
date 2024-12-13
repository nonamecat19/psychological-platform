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

func GetMyPrivateChatsHandler(c *fiber.Ctx) error {
	user := GetUserAuthInfo(c)
	privateChatsRepository := repository.NewPrivateChatRepository()

	specialistChats, _ := privateChatsRepository.FindAllBySpecialist(user.ID)
	clientChats, _ := privateChatsRepository.FindAllByClient(user.ID)

	combinedChats := append(specialistChats, clientChats...)

	return c.JSON(combinedChats)
}

func GetPrivateChatById(c *fiber.Ctx) error {
	intId, _ := strconv.Atoi(c.Params("id"))
	id := uint(intId)

	privateChatsRepository := repository.NewPrivateChatRepository()

	chat, err := privateChatsRepository.FindById(id)
	if err == nil {
		return c.JSON(chat[0])
	}
	return c.JSON(fiber.Map{"message": "Chat not found"})
}

func GetPrivateChatIdByUserId(c *fiber.Ctx) error {
	user := GetUserAuthInfo(c)
	intId, _ := strconv.Atoi(c.Params("id"))
	id := uint(intId)

	privateChatsRepository := repository.NewPrivateChatRepository()

	if user.Role == "user" {
		chat, err := privateChatsRepository.FindConcreteChat(user.ID, id)
		if err == nil {
			return c.JSON(chat.ID)
		}
		if err.Error() == "record not found" {
			chat, _ = privateChatsRepository.CreateConcreteChat(user.ID, id)
			return c.JSON(chat.ID)
		}
		panic(err)
	}
	if user.Role == "psychologist" {
		chat, err := privateChatsRepository.FindConcreteChat(id, user.ID)
		if err == nil {
			return c.JSON(chat.ID)
		}
		if err.Error() == "record not found" {
			chat, _ = privateChatsRepository.CreateConcreteChat(id, user.ID)
			return c.JSON(chat.ID)
		}
	}
	panic("Invalid role")
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
