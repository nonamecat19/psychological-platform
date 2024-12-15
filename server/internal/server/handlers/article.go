package handlers

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/model"
	"server/internal/repository"
	"strconv"
)

func GetAllArticlesHandler(c *fiber.Ctx) error {
	articlesRepository := repository.NewArticleRepository()
	articles, _ := articlesRepository.FindAll()

	return c.JSON(articles)
}

func CreateArticleHandler(c *fiber.Ctx) error {
	articlesRepository := repository.NewArticleRepository()

	data := new(model.Article)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	article, _ := articlesRepository.Create(*data)

	return c.JSON(article)
}

func UpdateArticleHandler(c *fiber.Ctx) error {
	articlesRepository := repository.NewArticleRepository()

	data := new(model.Article)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	article, _ := articlesRepository.UpdateOne(*data)

	return c.JSON(article)
}

func DeleteArticleHandler(c *fiber.Ctx) error {
	articlesRepository := repository.NewArticleRepository()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = articlesRepository.DeleteOne(uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
