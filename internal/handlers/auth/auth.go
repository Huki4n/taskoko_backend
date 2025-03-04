package auth

import (
	"awesomeProject/ent"
	"awesomeProject/ent/user"
	. "awesomeProject/internal/models"
	"awesomeProject/internal/service/refresh_token"
	"awesomeProject/pkg/errorsResponse"
	"awesomeProject/pkg/tokens"
	"context"
	_ "errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Handler содержит клиент Ent
type Handler struct {
	Client *ent.Client
}

// NewHandler создаёт новый обработчик с клиентом Ent
func NewHandler(client *ent.Client) *Handler {
	return &Handler{Client: client}
}

// LoginResponse - ответ при успешной авторизации
type LoginResponse struct {
	Status      string `json:"status"`
	AccessToken string `json:"accessToken"`
}

// Login логинит пользователя
func (h *Handler) Login(c echo.Context) error {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	u, err := h.Client.User.
		Query().
		Where(user.Email(request.Email)).
		Only(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: fmt.Sprintf("Not Found User by email %s", request.Email),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(request.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, Response{
			Status:  "Error",
			Message: "Wrong Password",
		})
	}

	accessToken, errAccess := tokens.GenerateAccessToken(u.Name)
	refreshToken, errRefresh := tokens.GenerateRefreshToken(u.Name)

	if errAccess != nil || errRefresh != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Failed to generate refresh_token",
		})
	}

	refresh_token.SetRefreshToken(c, refreshToken)

	return c.JSON(http.StatusOK, LoginResponse{
		Status:      "Success",
		AccessToken: accessToken,
	})
}

// Register регистрирует пользователя
func (h *Handler) Register(c echo.Context) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Failed to hash password",
		})
	}

	newUser, err := h.Client.User.Create().
		SetName(request.Username).
		SetEmail(request.Email).
		SetPassword(string(hashedPassword)).
		Save(context.Background())

	if err != nil {
		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusConflict, Response{
				Status:  "Error",
				Message: "User already exists",
			})
		}

		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, Response{
		Status:  "Success",
		Message: fmt.Sprintf("User %s has been registered", newUser.Name),
	})
}

// Refresh обновляет access token пользователя
func (h *Handler) Refresh(c echo.Context) error {

	refreshTokenClaims, err := refresh_token.GetClaims(c)
	if err != nil {
		return errorsResponse.HandleTokenError(c, err)
	}

	accessToken, errAccess := tokens.GenerateAccessToken(refreshTokenClaims.Username)
	if errAccess != nil {
		return errorsResponse.HandleTokenError(c, errAccess)
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Status:      "Success",
		AccessToken: accessToken,
	})
}
