package routes

import (
	"awesomeProject/ent"
	"awesomeProject/internal/routes/auth"
	"awesomeProject/internal/routes/project"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes регистрирует все маршруты API
func RegisterRoutes(e *echo.Echo, client *ent.Client) {
	api := e.Group("/api")

	auth.RegisterAuthRoutes(api, client)
	project.RegisterProjectRoutes(api, client)
}
