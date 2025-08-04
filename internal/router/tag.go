package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"zll.blog.com/internal/controller"
	"zll.blog.com/internal/pb"
)

func RegisterTagRouter(r *gin.Engine, conn *grpc.ClientConn) {
	tagGroup := r.Group("/tag")

	tagClient := pb.NewTagServiceClient(conn)
	tagController := controller.NewTagController(tagClient)
	{
		tagGroup.GET("/all", tagController.GetAll)
	}
}
