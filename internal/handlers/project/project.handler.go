package handler

import (
	"awesomeProject/internal/models"
	projectModels "awesomeProject/internal/models/project"
	service "awesomeProject/internal/service/project"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: service}
}

// GetProject обрабатывает POST /api/project/get-project/:id
func (h *ProjectHandler) GetProject(c echo.Context) error {
	id := c.Param("id")

	projectID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid project ID",
		})
	}

	getProject, err := h.projectService.GetProject(c.Request().Context(), projectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	// Возвращаем найденный проект
	return c.JSON(http.StatusOK, getProject)
}

// GetProjects обрабатывает POST /api/project/get-projects
func (h *ProjectHandler) GetProjects(c echo.Context) error {
	projects, err := h.projectService.GetProjects(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, projects)
}

// CreateProject обрабатывает POST /api/project/create-project
func (h *ProjectHandler) CreateProject(c echo.Context) error {
	var createProjectRequest projectModels.CreateProjectRequest

	if err := c.Bind(&createProjectRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, models.Response{
			Status:  "Error",
			Message: "User not authenticated",
		})
	}

	createProject, err := h.projectService.CreateProject(c.Request().Context(), createProjectRequest, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, createProject)
}

// EditProject обрабатывает PATCH /api/project/edit-project/:id
// func (h *Handler) EditProject(c echo.Context) error {
//	idParam := c.Param("id")
//	id, err := strconv.Atoi(idParam)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, models.Response{
//			Status:  "Error",
//			Message: "Invalid message ID",
//		})
//	}
//
//	var input struct {
//		Text string `json:"text"`
//	}
// }
//
// DeleteProject обрабатывает Delete /api/project/delete-project/:id
// func (h *Handler) DeleteProject(c echo.Context) error {
//
// }
