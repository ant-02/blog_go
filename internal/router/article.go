package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"zll.blog.com/internal/controller"
	"zll.blog.com/internal/pb"
)

func RegisterArticleRouter(r *gin.Engine, conn *grpc.ClientConn) {
	articleGroup := r.Group("/article")

	articleClient := pb.NewArticleServiceClient(conn)
	articleCategoryClient := pb.NewArticleCategoryServiceClient(conn)
	articleController := controller.NewArticleController(articleClient, articleCategoryClient)
	{
		articleGroup.GET("/:id", articleController.GetArticleById)
		articleGroup.GET("/dtos/categoryId/:id/:count", articleController.GetArticleDTOsByCategoryId)
		articleGroup.GET("/dtos/keywords/:keywords", articleController.GetArticleDTOsByKeywords)
		articleGroup.GET("/dtos/userId/:userId/:page/:pageSize/:status", articleController.GetArticleDTOsByUserId)
	}

	articleAuthGroup := articleGroup.Group("/auth")
	{
		articleAuthGroup.PUT("", articleController.SaveArticle)
		articleAuthGroup.POST("", articleController.SaveArticle)
	}
}
