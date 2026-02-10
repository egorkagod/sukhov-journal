package auth


import (
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(secret []byte, roles []string, userID *int64) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			claims, error := ParseAndValidateTokenFromHeader(secret, authHeader)
			if error != nil {
				return echo.NewHTTPError(401, "Токен неверный или отсутствует")
			}

			if roles != nil && hasRole(roles, claims.Role) {
				return next(c)
			}

			if userID != nil && claims.UserID == *userID {
				return next(c)
			}

			if roles == nil && userID == nil {
				return next(c)
			}

			return echo.NewHTTPError(403, "Недостаточно прав")
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