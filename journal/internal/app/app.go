package app

import (
	"github.com/labstack/echo/v4"
	"journal/internal/config"

	"journal/internal/articles"
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

func RegisterArticleApp(e *echo.Echo, config *config.Config) {
	articleRepo := articles.NewRepo()
	articleService := articles.NewService(articleRepo)
	articleHandler := articles.NewArticleHandler(articleService)
	articles := e.Group("/article", auth.AuthMiddleware(config.JWTSecret))
	articles.GET("/get", articleHandler.GetView)
	articles.POST("/create", articleHandler.CreateView)
	articles.PATCH("/edit", articleHandler.EditView)
}

func New(e *echo.Echo, config *config.Config) {
	RegisterAuthApp(e, config)
	RegisterArticleApp(e, config)
}
