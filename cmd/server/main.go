package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/db"
	"github.com/raphacosil/go-study/internal/handler"
	"github.com/raphacosil/go-study/internal/repository"
	"github.com/raphacosil/go-study/internal/service"
)

func main() {
	database := db.NewPostgresConnection()

	customerRepo := repository.NewCustomerRepository(database)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	r := gin.Default()
	r.GET("/users", customerHandler.FindAll)

	r.Run(":8080")
}