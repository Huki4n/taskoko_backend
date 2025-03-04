package tokens

import (
	errorsResponse "awesomeProject/pkg/errorsResponse"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(os.Getenv("SECRET_KEY"))
var RefreshSecretKey = []byte(os.Getenv("REFRESH_SECRET_KEY"))

var AccessTokenLiveTime = 10 * time.Second
var RefreshTokenLiveTime = 20 * time.Second

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateAccessToken создает JWT access токен
func GenerateAccessToken(username string) (accessToken string, err error) {
	accessToken, err = generateJWT(username, SecretKey, AccessTokenLiveTime)
	if err != nil {
		return "", errorsResponse.ErrTokenGenerationFailed
	}

	return accessToken, nil
}

// GenerateRefreshToken создает JWT access токен
func GenerateRefreshToken(username string) (refreshToken string, err error) {
	refreshToken, err = generateJWT(username, RefreshSecretKey, RefreshTokenLiveTime)
	if err != nil {
		return "", errorsResponse.ErrTokenGenerationFailed
	}

	return refreshToken, nil
}

// ValidateToken проверяет токен
func ValidateToken(tokenString string, isRefresh bool) (*Claims, error) {
	secret := SecretKey
	if isRefresh {
		secret = RefreshSecretKey
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		// Обработка различных типов ошибок
		if errors.Is(err, jwt.ErrTokenExpired) {
			fmt.Printf("JWT validation error: %v", err)
			return nil, errorsResponse.ErrTokenExpired
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errorsResponse.ErrTokenInvalid
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, errorsResponse.ErrTokenInvalid
		} else {
			// Логирование внутренней ошибки может быть полезным
			return nil, errorsResponse.ErrTokenInvalid
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errorsResponse.ErrTokenClaims
}

// generateJWT - функция для создания токенов
func generateJWT(username string, secretKey []byte, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
