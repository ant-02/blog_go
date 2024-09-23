package routers

import (
	"os"
	"time"

	_ "blog/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	r := gin.Default()
	
	// 解决跨域问题
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"Origin"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// 文章路由
	article(r)
	// 标签路由
	tag(r)
	// 文章标签路由
	articleTag(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(os.Getenv("SERVE_POST"))
}
