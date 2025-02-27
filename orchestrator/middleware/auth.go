package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/neandrson/go-daev4/orchestrator/utils"
)

func AuthenticateMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Отсутствует заголовок 'Authorization'",
		})
	}

	// Проверяем валидность и подпись токена
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return utils.Secret, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Ошибка при проверке токена: " + err.Error(),
		})
	}

	// Получаем данные пользователя из токена и сохраняем их в контексте запроса
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	c.Locals("email", email)

	return c.Next()
}
