package routes

import (
	"awesomeProject/ent"
	"awesomeProject/internal/handlers/message"
	"awesomeProject/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterMessageRoutes(api *echo.Group, client *ent.Client) {
	handlerMessage := message.NewHandler(client)

	messageGroup := api.Group("/message")

	messageGroup.GET("", handlerMessage.GetMessages, middleware.AuthMiddleware)
	messageGroup.POST("", handlerMessage.CreateMessage)
	messageGroup.PATCH("/:id", handlerMessage.UpdateMessage)
	messageGroup.DELETE("/:id", handlerMessage.DeleteMessage)
}
