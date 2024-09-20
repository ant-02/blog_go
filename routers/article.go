package routers

import (
	"blog/handlers"

	"github.com/gin-gonic/gin"
)

func article(r *gin.Engine) {
	getArticles(r)
	createArticle(r)
	getArticleById(r)
}

func getArticles(r *gin.Engine) {
	r.GET("/articles", handlers.GetArticles)
}

func createArticle(r *gin.Engine) {
	r.POST("/article", handlers.CreateArticle)
}

func getArticleById(r *gin.Engine) {
	r.GET("/article/:id", handlers.GetArticleById)
}