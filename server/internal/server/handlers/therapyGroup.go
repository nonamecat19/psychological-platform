package handlers

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/model"
	"server/internal/repository"
	"strconv"
)

func GetAllTherapyGroupsHandler(c *fiber.Ctx) error {
	therapyGroupsRepository := repository.NewTherapyGroupRepository()
	groups, _ := therapyGroupsRepository.FindAll()

	return c.JSON(groups)
}

func GetTherapyGroupByIdHandler(c *fiber.Ctx) error {
	intId, _ := strconv.Atoi(c.Params("id"))
	id := uint(intId)

	therapyGroupsRepository := repository.NewTherapyGroupRepository()

	group, err := therapyGroupsRepository.FindById(id)
	if err == nil {
		return c.JSON(group[0])
	}
	return c.JSON(fiber.Map{"message": "Group not found"})
}

func GetAllMyTherapyGroupsHandler(c *fiber.Ctx) error {
	user := GetUserAuthInfo(c)

	therapyGroupsRepository := repository.NewTherapyGroupRepository()

	if user.Role == "psychologist" {
		groups, _ := therapyGroupsRepository.FindAllBySpecialist(user.ID)
		return c.JSON(groups)
	}

	groups, _ := therapyGroupsRepository.FindAll()
	return c.JSON(groups)
}

func CreateTherapyGroupHandler(c *fiber.Ctx) error {
	user := GetUserAuthInfo(c)

	therapyGroupsRepository := repository.NewTherapyGroupRepository()

	data := new(model.TherapyGroup)

	data.SpecialistID = user.ID

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	group, _ := therapyGroupsRepository.Create(*data)

	return c.JSON(group)
}

func UpdateTherapyGroupHandler(c *fiber.Ctx) error {
	therapyGroupsRepository := repository.NewTherapyGroupRepository()

	data := new(model.TherapyGroup)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	group, _ := therapyGroupsRepository.UpdateOne(*data)

	return c.JSON(group)
}

func DeleteTherapyGroupHandler(c *fiber.Ctx) error {
	therapyGroupsRepository := repository.NewTherapyGroupRepository()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = therapyGroupsRepository.DeleteOne(uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
