package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/db"
	"github.com/raphacosil/go-study/internal/handler"
	"github.com/raphacosil/go-study/internal/repository"
	"github.com/raphacosil/go-study/internal/service"
)

func main() {
	database := db.NewMySqlConnection()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	r.GET("/users", userHandler.FindAll)

	r.Run(":8080")
}