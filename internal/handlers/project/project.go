package message

import (
	"awesomeProject/ent"
	"awesomeProject/ent/project"
	m "awesomeProject/internal/models"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"log"
	"net/http"
	"strconv"
)

type GetProjectResponse struct {
	ID      int                  `json:"id"`
	Name    string               `json:"name"`
	Users   []ProjectUser        `json:"users"`
	Columns []*ent.ProjectColumn `json:"columns"`
}

type GetProjectsResponse struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Users []ProjectUser `json:"users"`
}

type ProjectUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Handler содержит клиент Ent
type Handler struct {
	Client *ent.Client
}

// NewProjectHandler создаёт новый обработчик с клиентом Ent
func NewProjectHandler(client *ent.Client) *Handler {
	return &Handler{Client: client}
}

// GetProject обрабатывает POST /api/project/get-project/:id
func (h *Handler) GetProject(c echo.Context) error {
	id := c.Param("id")

	projectID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, m.Response{
			Status:  "Error",
			Message: "Invalid project ID",
		})
	}

	// Ищем проект по ID
	getProject, err := h.Client.Project.Query().
		Where(project.ID(projectID)).
		WithColumns().
		WithUsers(
			func(q *ent.UserQuery) {
				q.Select("id", "name", "email")
			}).
		Only(c.Request().Context())

	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, m.Response{
				Status:  "Error",
				Message: "Project not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, m.Response{
			Status:  "Error",
			Message: "Failed to retrieve project",
		})
	}

	response := GetProjectResponse{
		ID:      getProject.ID,
		Name:    getProject.Name,
		Columns: getProject.Edges.Columns,
	}

	// Возвращаем найденный проект
	return c.JSON(http.StatusOK, response)
}

// GetProjects обрабатывает POST /api/project/get-projects
func (h *Handler) GetProjects(c echo.Context) error {
	projects, err := h.Client.Project.Query().WithUsers(
		func(q *ent.UserQuery) {
			q.Select("id", "name", "email")
		}).All(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, m.Response{
			Status:  "Error",
			Message: "Failed to retrieve projects",
		})
	}

	var response []GetProjectsResponse
	for _, p := range projects {
		response = append(response, GetProjectsResponse{
			ID:   p.ID,
			Name: p.Name,
			Users: lo.Map(p.Edges.Users, func(user *ent.User, _ int) ProjectUser {
				return ProjectUser{
					ID:    user.ID,
					Name:  user.Name,
					Email: user.Email,
				}
			}),
		})
	}

	// Возвращаем найденные проекты
	return c.JSON(http.StatusOK, response)
}

// CreateProject обрабатывает POST /api/project/create-project
func (h *Handler) CreateProject(c echo.Context) error {
	var createProjectRequest struct {
		ProjectName string `json:"projectName"`
	}

	if err := c.Bind(&createProjectRequest); err != nil {
		return c.JSON(http.StatusBadRequest, m.Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	tx, err := h.Client.Tx(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, m.Response{
			Status:  "Error",
			Message: "Transaction failed",
		})
	}

	createProject, err := tx.Project.Create().SetName(createProjectRequest.ProjectName).Save(context.Background())
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("transaction rollback failed: %v", rbErr)
		}

		return c.JSON(http.StatusInternalServerError, m.Response{
			Status:  "Error",
			Message: "Could not create the project",
		})
	}

	_, err = tx.ProjectColumn.CreateBulk(
		tx.ProjectColumn.Create().SetName("To Do").SetColor("#9d66d5").SetOrder(1).SetProject(createProject),
		tx.ProjectColumn.Create().SetName("In Work").SetColor("#3361ff").SetOrder(2).SetProject(createProject),
		tx.ProjectColumn.Create().SetName("Review").SetColor("#ffcd43").SetOrder(3).SetProject(createProject),
		tx.ProjectColumn.Create().SetName("Done").SetColor("#98db7c").SetOrder(4).SetProject(createProject),
	).Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, m.Response{
			Status:  "Error",
			Message: "Could not create the default columns",
		})
	}

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, m.Response{
			Status:  "Error",
			Message: "User not authenticated",
		})
	}

	_, err = tx.Project.Update().
		Where(project.ID(createProject.ID)).
		AddUserIDs(userID).
		Save(context.Background())
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("transaction rollback failed: %v", rbErr)
		}

		return c.JSON(http.StatusInternalServerError, m.Response{
			Status:  "Error",
			Message: "Could not add user to the project",
		})
	}

	if commitErr := tx.Commit(); commitErr != nil {
		log.Printf("transaction commit failed: %v", commitErr)
	}

	return c.JSON(http.StatusOK, createProject)
}

// EditProject обрабатывает PATCH /api/project/edit-project/:id
//func (h *Handler) EditProject(c echo.Context) error {
//	idParam := c.Param("id")
//	id, err := strconv.Atoi(idParam)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, m.Response{
//			Status:  "Error",
//			Message: "Invalid message ID",
//		})
//	}
//
//	var input struct {
//		Text string `json:"text"`
//	}
//}
//
//// DeleteProject обрабатывает Delete /api/project/delete-project/:id
//func (h *Handler) DeleteProject(c echo.Context) error {
//
//}
