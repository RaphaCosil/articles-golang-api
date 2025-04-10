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
		customerGroup.GET("/:id", customerHandler.FindById)
		customerGroup.POST("", customerHandler.Create)
		customerGroup.PUT("/:id", customerHandler.Update)
		customerGroup.DELETE("/:id", customerHandler.DeleteById)
	}

	articleGroup := r.Group("/article")
	{
		articleGroup.GET("", articleHandler.FindAll)
		articleGroup.GET("/:id", articleHandler.FindById)
		articleGroup.POST("", articleHandler.Create)
		articleGroup.PUT("/:id", articleHandler.Update)
		articleGroup.DELETE("/:id", articleHandler.DeleteById)
		articleGroup.GET("/customer/:customerId", articleHandler.FindByCustomerId)
		articleGroup.GET("/title/:title", articleHandler.FindByTitle)
		articleGroup.GET("/content/:content", articleHandler.FindByContent)
		articleGroup.GET("/keywords/:keywords", articleHandler.FindByKeywords)
		articleGroup.GET("/keywords/filter/:keywords", articleHandler.FindByKeywordsFilter)
	}

	keywordGroup := r.Group("/keyword")
	{
		keywordGroup.GET("", keywordHandler.FindAll)
		keywordGroup.GET("/:id", keywordHandler.FindById)
		keywordGroup.POST("", keywordHandler.Create)
		keywordGroup.PUT("/:id", keywordHandler.Update)
		keywordGroup.DELETE("/:id", keywordHandler.DeleteById)
		keywordGroup.GET("/article/:articleId", keywordHandler.FindByArticleId)
	}

	return r
}
