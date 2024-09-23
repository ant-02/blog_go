package routers

import (
	"blog/internal/handlers"

	"github.com/gin-gonic/gin"
)

func article(r *gin.Engine) {
	createArticle(r)
	getArticleById(r)
}


func createArticle(r *gin.Engine) {
	r.POST("/article", handlers.CreateArticle)
}

func getArticleById(r *gin.Engine) {
	r.GET("/article/:id", handlers.GetArticleById)
}