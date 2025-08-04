package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"zll.blog.com/internal/controller"
	"zll.blog.com/internal/pb"
)

func RegisterArticleCategoryRouter(r *gin.Engine, conn *grpc.ClientConn) {
	articleCategoryGroup := r.Group("/articleCategory")

	articleCategoryClient := pb.NewArticleCategoryServiceClient(conn)
	articleCategoryController := controller.NewArticleCategoryController(articleCategoryClient)
	{
		articleCategoryGroup.GET("/categoryId/:id", articleCategoryController.GetCategoryIdByArticleId)
	}
	articleCategoryAuthGroup := articleCategoryGroup.Group("/auth")
	{
		articleCategoryAuthGroup.PUT("", articleCategoryController.SaveArticleCategory)
	}
}