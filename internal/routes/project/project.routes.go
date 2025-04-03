package project

import (
	"awesomeProject/ent"
	"awesomeProject/internal/handlers/project"
	"awesomeProject/internal/middleware"
	"awesomeProject/internal/repository/project"
	"awesomeProject/internal/service/project"
	"github.com/labstack/echo/v4"
)

func RegisterProjectRoutes(api *echo.Group, client *ent.Client) {
	repo := repository.NewProjectRepository(client)
	srv := service.NewProjectService(repo)
	handlerProject := handler.NewProjectHandler(srv)

	projectGroup := api.Group("/project", middleware.AuthMiddleware)

	projectGroup.POST("/create-project", handlerProject.CreateProject)
	projectGroup.GET("/get-project/:id", handlerProject.GetProject)
	projectGroup.GET("/get-projects", handlerProject.GetProjects)
	projectGroup.DELETE("/delete-project/:id", handlerProject.DeleteProject)
}
