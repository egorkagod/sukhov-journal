package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"journal/internal/articles"
	"journal/internal/auth"
	"journal/internal/config"
	"journal/internal/db"
	"journal/internal/voice"
)

func InitDB(config *config.Config) *gorm.DB {
	credentials := &db.DBCredentials{Host: config.DBHost, User: config.DBUser, Password: config.DBPassword, Name: config.DBName, Port: config.DBPort}
	db, err := db.NewDB(credentials)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&auth.User{}, &articles.Article{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func RegisterAuthApp(e *echo.Echo, config *config.Config, db *gorm.DB) {
	authRepo, err := auth.NewRepo(db)
	if err != nil {
		log.Fatalf("Ошибка при создании репозитория - %v", err.Error())
	}

	authService := auth.NewService(authRepo)
	authHandler := auth.NewAuthHandler(config.JWTSecret, authService)
	e.GET("/user", authHandler.GetView)
	e.POST("/login", authHandler.LoginView)
	e.POST("/register", authHandler.RegisterView)
	e.POST("/logout", authHandler.LogoutView)
}

func RegisterArticleApp(e *echo.Echo, config *config.Config, db *gorm.DB) {
	articleRepo := articles.NewRepo(db)
	articleService := articles.NewService(articleRepo)
	articleHandler := articles.NewArticleHandler(articleService)
	articles := e.Group("/article", auth.AuthMiddleware(config.JWTSecret))
	articles.GET("/get", articleHandler.GetView)
	articles.POST("/create", articleHandler.CreateView)
	articles.PATCH("/edit", articleHandler.EditView)
	articles.DELETE("/delete", articleHandler.DeleteView)
	articles.GET("/speech", articleHandler.VoiceOverView)
}

func New(e *echo.Echo, config *config.Config) {
	voice.InitManager(config)
	db := InitDB(config)
	RegisterAuthApp(e, config, db)
	RegisterArticleApp(e, config, db)
}
