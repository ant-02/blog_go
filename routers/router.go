package routers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	// 文章路由
	article(r)

	r.Run(os.Getenv("SERVE_POST"))
}