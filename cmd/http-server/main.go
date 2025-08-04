package main

import (
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "zll.blog.com/docs"
	"zll.blog.com/internal/config"
	"zll.blog.com/internal/router"
)

// @title blog
// @version 1.0
// @description 这是一个博客项目
// @host localhost:8080
// @BasePath /
func main() {
	r := router.RegisterAllRouters()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":" + config.GetConfig().Server.Port); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
