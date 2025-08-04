package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"zll.blog.com/internal/controller"
	"zll.blog.com/internal/pb"
)

func RegisterArticleTagRouter(r *gin.Engine, conn *grpc.ClientConn) {
	articleTagGroup := r.Group("/articleTag")

	articleTagClient := pb.NewArticleTagServiceClient(conn)
	articleTagController := controller.NewArticleTagController(articleTagClient)
	{
		articleTagGroup.GET("/ids/:id", articleTagController.GetTagIdsByArticleId)
	}
}