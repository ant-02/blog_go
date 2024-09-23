package routers

import (
	"blog/internal/handlers"

	"github.com/gin-gonic/gin"
)

func articleTag(r *gin.Engine) {
	getArticlePreviewTags(r)
}

func getArticlePreviewTags(r *gin.Engine) {
	r.GET("/articleTags", handlers.GetArticlePreviewTags)
}