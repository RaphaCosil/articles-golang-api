package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/handler"
)

func SetupRouter(
	customerHandler *handler.CustomerHandler,
	articleHandler *handler.ArticleHandler,
	keywordHandler *handler.KeywordHandler,
	) *gin.Engine {
	r := gin.Default()

	customerGroup := r.Group("/customer")
	{
		customerGroup.GET("", customerHandler.FindAll)
		customerGroup.GET("/:id", customerHandler)
		customerGroup.POST("", customerHandler.CreateUser)
		customerGroup.PUT("/:id", customerHandler.UpdateUser)
		customerGroup.DELETE("/:id", customerHandler.DeleteUser)
	}

	articleGroup := r.Group("/article")
	{
		articleGroup.GET("", articleHandler.FindAll)
		articleGroup.GET("/:id", articleHandler)
		articleGroup.POST("", articleHandler.CreateUser)
		articleGroup.PUT("/:id", articleHandler.UpdateUser)
		articleGroup.DELETE("/:id", articleHandler.DeleteUser)
	}

	keywordGroup := r.Group("/article")
	{
		keywordGroup.GET("", keywordHandler.FindAll)
		keywordGroup.GET("/:id", keywordHandler)
		keywordGroup.POST("", keywordHandler.CreateUser)
		keywordGroup.PUT("/:id", keywordHandler.UpdateUser)
		keywordGroup.DELETE("/:id", keywordHandler.DeleteUser)
	}

	return r
}
