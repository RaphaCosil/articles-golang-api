package main

import (
	"github.com/raphacosil/go-study/internal/db"
	"github.com/raphacosil/go-study/internal/handler"
	"github.com/raphacosil/go-study/internal/repository"
	"github.com/raphacosil/go-study/internal/service"
	"github.com/raphacosil/go-study/internal/router"
)

func main() {
	database := db.NewPostgresConnection()

	customerRepo := repository.NewCustomerRepository(database)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	articleRepo := repository.NewArticleRepository(database)
	articleService := service.NewArticleService(articleRepo)
	articleHandler := handler.NewArticleHandler(articleService)

	keywordRepo := repository.NewKeywordRepository(database)
	keywordService := service.NewKeywordService(keywordRepo)
	keywordHandler := handler.NewKeywordHandler(keywordService)

	r := router.SetupRouter(
		customerHandler,
		articleHandler,
		keywordHandler,
	)

	r.Run(":8080")
}