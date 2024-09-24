package routers

import (
	"blog/internal/handlers"

	"github.com/gin-gonic/gin"
)

func articleTag(r *gin.Engine) {
	getArticlePreviewTags(r)
	getArticlePreviewTag(r)
}

func getArticlePreviewTags(r *gin.Engine) {
	r.GET("/articleTags", handlers.GetArticlePreviewTags)
}

func getArticlePreviewTag(r *gin.Engine) {
	r.GET("/articleTag/:id", handlers.GetArticlePreviewTag)
}