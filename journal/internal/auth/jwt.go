package auth

import (
	"time"
	"errors"
	"strings"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64 `json:"uid"`
	Role string `json:"role"`
	jwt.RegisteredClaims 
}

func GenerateJWT(secret []byte, userID int64, role string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseAndValidateJWT(secret []byte, tokenStr string) (*Claims, error) {
	claims := &Claims{}
	
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("Неизвестный алгоритм подписи")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid{
		return nil, errors.New("Неверный токен")
	}

	return claims, nil
}

func ParseAndValidateTokenFromHeader(secret []byte, authHeader string) (*Claims, error) {
	const prefix = "Bearer "
	
	if !strings.HasPrefix(authHeader, prefix) {
		return nil, errors.New("Отсутствует заголовок Bearer")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	
	if token == "" {
		return nil, errors.New("Пустой токен")
	}

	claims, err := ParseAndValidateJWT(secret, token)

	if err != nil {
		return nil, err
	}

	return claims, nil
}