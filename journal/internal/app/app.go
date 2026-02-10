package app

import (
	"journal/internal/config"
	"github.com/labstack/echo/v4"

	"journal/internal/auth"
)

func RegisterAuthApp(e *echo.Echo, config *config.Config) {
	authRepo := auth.NewRepo()
	authService := auth.NewService(authRepo)
	authHandler := auth.NewAuthHandler(config.JWTSecret, authService)
	e.POST("/login", authHandler.LoginView)
	e.POST("/register", authHandler.RegisterView)
	e.POST("/logout", authHandler.LogoutView)
}

// func RegisterArticleApp(e *echo.Echo, config *config.Config) {}

func New(e *echo.Echo, config *config.Config) {
	RegisterAuthApp(e, config)
	// Article app
}