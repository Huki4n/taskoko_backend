package message

//
//import (
//	"awesomeProject/ent"
//	m "awesomeProject/internal/models"
//	"context"
//	"github.com/labstack/echo/v4"
//	"net/http"
//	"strconv"
//)
//
//// Handler содержит клиент Ent
//type Handler struct {
//	Client *ent.Client
//}
//
//// NewTaskHandler создаёт новый обработчик с клиентом Ent
//func NewTaskHandler(client *ent.Client) *Handler {
//	return &Handler{Client: client}
//}
//
//// CreateTask обрабатывает POST /api/task/create-task
//func (h *Handler) CreateTask(c echo.Context) error {
//	var input struct {
//		Text string `json:"text"`
//	}
//
//	if err := c.Bind(&input); err != nil {
//		return c.JSON(http.StatusBadRequest, m.Response{
//			Status:  "Error",
//			Message: "Invalid input",
//		})
//	}
//
//	msg, err := h.Client.Task.Create()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, m.Response{
//			Status:  "Error",
//			Message: "Could not add the message",
//		})
//	}
//
//	return c.JSON(http.StatusOK, msg)
//}
//
//// EditTask обрабатывает PATCH /api/task/edit-task
//func (h *Handler) EditTask(c echo.Context) error {
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
//
//}
//
//// DeleteTask обрабатывает Delete /api/task/delete-task
//func (h *Handler) DeleteTask(c echo.Context) error {
//
//}
