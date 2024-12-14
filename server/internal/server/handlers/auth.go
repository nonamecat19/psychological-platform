package handlers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"log"
	lib "server/internal/lib"
	"server/internal/model"
	"server/internal/repository"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	IsAnonymous bool   `json:"isAnonymous"`
	Role        string `json:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func AuthLoginHandler(c *fiber.Ctx) error {
	userRepository := repository.NewUserRepository()

	loginRequest := new(LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := userRepository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(LoginResponse{
		Token: t,
	})
}

func AuthRegisterHandler(c *fiber.Ctx) error {
	userRepository := repository.NewUserRepository()

	registerRequest := new(RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newUser := model.User{
		Name:        &registerRequest.Name,
		Surname:     &registerRequest.Surname,
		IsAnonymous: registerRequest.IsAnonymous,
		Role:        registerRequest.Role,
		Email:       registerRequest.Email,
		Password:    registerRequest.Password,
	}

	user, err := userRepository.Create(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(LoginResponse{
		Token: t,
	})
}

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or malformed JWT"})
	}

	claims, err := lib.GetClaimsFromJWT(authHeader)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository()
	userId := claims["ID"]
	user, err := userRepository.GetOne(uint(userId.(float64)))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT User"})
	}

	c.Locals("user", claims)
	c.Locals("user-info", user)
	return c.Next()
}

func GetUserAuthInfo(c *fiber.Ctx) model.User {
	return c.Locals("user-info").(model.User)
}
