package middleware

import (
	. "awesomeProject/internal/models"
	"awesomeProject/pkg/tokens"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Извлекаем access_token из заголовка Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, Response{
				Status:  "Error",
				Message: "Authorization header is missing",
			})
		}

		// Проверяем, что заголовок содержит Bearer токен
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, Response{
				Status:  "Error",
				Message: "Invalid Authorization header format",
			})
		}

		accessToken := tokenParts[1]

		// Проверяем access_token
		_, err := tokens.ValidateToken(accessToken, false)
		if err != nil { // Если access_token недействителен
			return c.JSON(http.StatusUnauthorized, Response{
				Status:  "Error jwt token",
				Message: "Jwt token is invalid",
			})
		}

		// Если access_token валиден, продолжаем выполнение
		return next(c)
	}
}
