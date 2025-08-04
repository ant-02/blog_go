package router

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterAllRouters() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	RegisterArticleRouter(r, conn)
	RegisterCategoryRouter(r, conn)
	RegisterUserRouter(r, conn)
	RegisterUserFollowRouter(r, conn)
	RegisterTagRouter(r, conn)
	RegisterArticleTagRouter(r, conn)
	RegisterArticleCategoryRouter(r, conn)

	return r
}
