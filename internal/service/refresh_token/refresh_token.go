package refresh_token

import (
	"awesomeProject/pkg/errorsResponse"
	"awesomeProject/pkg/tokens"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SetRefreshToken(c echo.Context, refreshToken string) {
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(tokens.RefreshTokenLiveTime * 2),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
}

func RemoveRefreshToken(c echo.Context, refreshToken string) error {
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(0),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
	return nil
}

func GetClaims(c echo.Context) (*tokens.Claims, error) {
	refreshCookie, err := c.Cookie("refresh_token")
	if err != nil {
		return nil, errorsResponse.ErrMissingToken
	}

	refreshClaims, err := tokens.ValidateToken(refreshCookie.Value, true)
	if err != nil {
		return nil, err
	}

	return refreshClaims, nil
}
