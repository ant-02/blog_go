package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"zll.blog.com/internal/controller"
	"zll.blog.com/internal/pb"
)

func RegisterCategoryRouter(r *gin.Engine, conn *grpc.ClientConn) {
	categoryGroup := r.Group("/category")
	
	categoryClient := pb.NewCategoryServiceClient(conn)
	categoryController := controller.NewCategoryController(categoryClient)
	{
		categoryGroup.GET("/all", categoryController.GetAll)
	}
}
