package routers

import (
	"os"

	_ "blog/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	r := gin.Default()

	// 文章路由
	article(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(os.Getenv("SERVE_POST"))
}
