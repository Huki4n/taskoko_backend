package routes

import (
	"awesomeProject/ent"
	project "awesomeProject/internal/handlers/project"
	"awesomeProject/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterProjectRoutes(api *echo.Group, client *ent.Client) {
	handlerProject := project.NewProjectHandler(client)

	projectGroup := api.Group("/project", middleware.AuthMiddleware)

	projectGroup.POST("/create-project", handlerProject.CreateProject)
	projectGroup.GET("/get-project/:id", handlerProject.GetProject)
	projectGroup.GET("/get-projects", handlerProject.GetProjects)
}
