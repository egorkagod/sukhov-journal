package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	JWTSecret []byte
	svc       UserService
}

func NewAuthHandler(JWTSecret []byte, svc UserService) *AuthHandler {
	return &AuthHandler{JWTSecret: JWTSecret, svc: svc}
}

func (h *AuthHandler) GetView(c echo.Context) error {
	login := c.QueryParam("login")

	ctx := c.Request().Context()
	user, err := h.svc.GetRepo().GetByLogin(ctx, login)
	switch err {
	case UserNotFoundErr:
		return c.JSON(404, map[string]string{"message": "Пользователь не найден"})
	case nil:
	default:
		return echo.NewHTTPError(500, err)
	}
	return c.JSON(200, user)
}

func (h *AuthHandler) LoginView(c echo.Context) error {
	var data CredentialsSchema

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	if data.Login == "" || data.Password == "" {
		return echo.NewHTTPError(400, "Не передан логин или пароль")
	}

	ctx := c.Request().Context()
	dto := UserCredentialsServiceDTO{Login: data.Login, Password: data.Password}
	user, err := h.svc.Login(ctx, dto)
	switch err {
	case UserNotFoundErr:
		return echo.NewHTTPError(401, "Неверный логин или пароль")
	case nil:
	default:
		return echo.NewHTTPError(500, err.Error())
	}

	ttl := 15 * time.Minute

	token, err := GenerateJWT(h.JWTSecret, user.ID, user.Role, ttl)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    token,
		HttpOnly: true, // отключить при тесте
		MaxAge:   int(ttl.Seconds()),
	})

	return c.JSON(200, map[string]string{"message": "Успешный вход"})
}

func (h *AuthHandler) RegisterView(c echo.Context) error {
	var data CredentialsSchema

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	ctx := c.Request().Context()
	dto := UserCredentialsServiceDTO{Login: data.Login, Password: data.Password}
	err := h.svc.Register(ctx, dto)
	switch err {
	case nil:
		return c.JSON(200, map[string]string{"message": "Успешная регистрация"})
	case UserAlreadyExistsErr:
		return echo.NewHTTPError(409, UserAlreadyExistsErr.Error())
	default:
		return echo.NewHTTPError(500, err.Error())
	}
}

func (h AuthHandler) LogoutView(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   "access_token",
		Value:  "",
		MaxAge: -1,
	})

	return c.JSON(200, map[string]string{"message": "Успешный выход"})
}
