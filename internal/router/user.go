package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"zll.blog.com/internal/config"
	"zll.blog.com/internal/controller"
	middleware "zll.blog.com/internal/middleWare"
	"zll.blog.com/internal/pb"
)

func RegisterUserRouter(r *gin.Engine, conn *grpc.ClientConn) {
	userGroup := r.Group("/user")

	userClient := pb.NewUserServiceClient(conn)
	userController := controller.NewUserController(userClient)
	{
		userGroup.GET("/dto/:id", userController.GetUserDTOById)
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/dtos/:keywords", userController.GetUserDTOsByKeywords)
	}
	userAuthGroup := userGroup.Group("/auth")
	userAuthGroup.Use(middleware.JWTAuthMiddleWare(config.GetConfig().Jwt.SecretKey))
	{
		userAuthGroup.GET("/info", userController.GetUserById)
	}
}
