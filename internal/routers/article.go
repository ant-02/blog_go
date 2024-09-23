package routers

import (
	"blog/internal/handlers"

	"github.com/gin-gonic/gin"
)

func article(r *gin.Engine) {
	getArticlePreViews(r)
	createArticle(r)
	getArticleById(r)
}

func getArticlePreViews(r *gin.Engine) {
	r.GET("/articles", handlers.GetArticlePreview)
}

func createArticle(r *gin.Engine) {
	r.POST("/article", handlers.CreateArticle)
}

func getArticleById(r *gin.Engine) {
	r.GET("/article/:id", handlers.GetArticleById)
}