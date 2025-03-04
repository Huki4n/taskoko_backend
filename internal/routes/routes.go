package routes

import (
	"awesomeProject/ent"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes регистрирует все маршруты API
func RegisterRoutes(e *echo.Echo, client *ent.Client) {
	api := e.Group("/api")

	RegisterMessageRoutes(api, client)
	RegisterAuthRoutes(api, client)
}
