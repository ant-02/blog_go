package routers

import (
	"blog/internal/handlers"

	"github.com/gin-gonic/gin"
)

func tag(r *gin.Engine) {
	createTag(r)
}

func createTag(r *gin.Engine) {
	r.POST("/tag", handlers.CreateTag)
}