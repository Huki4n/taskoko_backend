package errorsResponse

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  string `json:"code"`
	Message string `json:"message"`
	Code    int    `json:"status"`
}

// HandleTokenError обрабатывает ошибки и возвращает соответствующий HTTP-ответ
func HandleTokenError(c echo.Context, err error) error {
	var errorResponse *ErrorResponse
	if errors.As(err, &errorResponse) {
		// Используем статус и данные из ошибки ErrorResponse
		return c.JSON(errorResponse.Code, Response{
			Status:  errorResponse.Status,
			Message: errorResponse.Message,
		})
	}

	// Для других типов ошибок используем статус 500
	return c.JSON(http.StatusInternalServerError, Response{
		Status:  "SERVER_ERROR",
		Message: "Internal server error",
	})
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

// Предопределённые ошибки валидации JWT-токенов
var (
	// ErrTokenExpired возникает, когда срок действия токена истек
	ErrTokenExpired = &ErrorResponse{
		Status:  "TOKEN_EXPIRED",
		Message: "Token is expired",
		Code:    401,
	}

	// ErrTokenInvalid возникает при неверном формате токена или подписи
	ErrTokenInvalid = &ErrorResponse{
		Status:  "INVALID_TOKEN",
		Message: "Invalid token format or signature",
		Code:    401,
	}

	// ErrTokenClaims возникает при проблемах с данными в токене
	ErrTokenClaims = &ErrorResponse{
		Status:  "INVALID_CLAIMS",
		Message: "Invalid token claims",
		Code:    400,
	}

	// ErrServerInternal возникает при внутренних ошибках сервера
	ErrServerInternal = &ErrorResponse{
		Status:  "SERVER_ERROR",
		Message: "Internal server error during token validation",
		Code:    500,
	}

	// ErrInvalidAlgorithm возникает при неподдерживаемом алгоритме подписи
	ErrInvalidAlgorithm = &ErrorResponse{
		Status:  "INVALID_ALGORITHM",
		Message: "Unexpected signing method",
		Code:    401,
	}

	// ErrMissingToken возникает, когда токен отсутствует в запросе
	ErrMissingToken = &ErrorResponse{
		Status:  "MISSING_TOKEN",
		Message: "Authentication token is missing",
		Code:    401,
	}

	// ErrTokenGenerationFailed возникает при ошибке создания токена
	ErrTokenGenerationFailed = &ErrorResponse{
		Status:  "TOKEN_GENERATION_FAILED",
		Message: "Failed to generate token",
		Code:    500,
	}
)
