package auth

import (
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(secret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("access_token")
			if err != nil {
				return echo.NewHTTPError(401, "Токен неверный или отсутствует")
			}

			claims, err := ParseAndValidateJWT(secret, cookie.Value)

			if err != nil {
				return echo.NewHTTPError(401, "Токен неверный или отсутствует")
			}
			c.Set("UserID", claims.UserID)

			return next(c)
		}
	}
}

func hasRole(roles []string, role string) bool {
	for i := 0; i < len(roles); i++ {
		if role == roles[i] {
			return true
		}
	}
	return false
}
