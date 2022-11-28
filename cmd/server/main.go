package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ibnumei/go-ms-playground/internal/app/database"
	"github.com/ibnumei/go-ms-playground/internal/app/user/handler"
	"github.com/ibnumei/go-ms-playground/internal/app/user/repository"
	"github.com/ibnumei/go-ms-playground/internal/app/user/service"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	route := gin.Default()

	db := database.DatabaseConnection()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	route.POST("/register", userHandler.Register)
	route.POST("/login", userHandler.Login)

	route.Run(":" + os.Getenv("PORT"))
}