package message

import (
	"context"
	"net/http"
	"strconv"

	"awesomeProject/ent"
	"github.com/labstack/echo/v4"
)

// Handler содержит клиент Ent
type Handler struct {
	Client *ent.Client
}

// NewHandler создаёт новый обработчик с клиентом Ent
func NewHandler(client *ent.Client) *Handler {
	return &Handler{Client: client}
}

// Response - стандартный формат ответа
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// GetMessages обрабатывает GET /api/message
func (h *Handler) GetMessages(c echo.Context) error {
	messages, err := h.Client.Message.Query().All(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Failed to retrieve messages",
		})
	}
	return c.JSON(http.StatusOK, messages)
}

// CreateMessage обрабатывает POST /api/message
func (h *Handler) CreateMessage(c echo.Context) error {
	var input struct {
		Text string `json:"text"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	msg, err := h.Client.Message.Create().SetText(input.Text).Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	return c.JSON(http.StatusOK, msg)
}

// UpdateMessage обрабатывает PATCH /api/message/:id
func (h *Handler) UpdateMessage(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid message ID",
		})
	}

	var input struct {
		Text string `json:"text"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid input JSON",
		})
	}

	msg, err := h.Client.Message.UpdateOneID(id).SetText(input.Text).Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	return c.JSON(http.StatusOK, msg)
}

// DeleteMessage обрабатывает DELETE /api/message/:id
func (h *Handler) DeleteMessage(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid message ID",
		})
	}

	err = h.Client.Message.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was deleted successfully",
	})
}
