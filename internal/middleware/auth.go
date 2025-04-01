package middleware

import (
	"net/http"
	"strings"

	"awesomeProject/internal/models"
	"awesomeProject/pkg/tokens"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Извлекаем access_token из заголовка Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Status:  "Error",
				Message: "Authorization header is missing",
			})
		}

		// Проверяем, что заголовок содержит Bearer токен
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Status:  "Error",
				Message: "Invalid Authorization header format",
			})
		}

		accessToken := tokenParts[1]

		// Проверяем access_token
		claims, err := tokens.ValidateToken(accessToken, false)
		if err != nil { // Если access_token недействителен
			return c.JSON(http.StatusUnauthorized, models.Response{
				Status:  "Error jwt token",
				Message: "Jwt token is invalid",
			})
		}

		userId := claims.UserId
		c.Set("userID", userId)

		// Если access_token валиден, продолжаем выполнение
		return next(c)
	}
}
