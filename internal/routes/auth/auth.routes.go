package auth

import (
	"awesomeProject/ent"
	"awesomeProject/internal/handlers/auth"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(api *echo.Group, client *ent.Client) {
	handlerAuth := auth.NewHandler(client)

	authGroup := api.Group("/auth")

	authGroup.POST("/login", handlerAuth.Login)
	authGroup.POST("/register", handlerAuth.Register)
	authGroup.GET("/refresh", handlerAuth.Refresh)
}
