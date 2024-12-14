package licurrentChatStoreb

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"strings"
)

func GetClaimsFromJWT(bearerToken string) (jtoken.MapClaims, error) {
	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")
	token, err := jtoken.Parse(tokenString, func(token *jtoken.Token) (interface{}, error) {
		if _, ok := token.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired JWT")
	}

	claims, ok := token.Claims.(jtoken.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
