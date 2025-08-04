package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"zll.blog.com/internal/config"
	"zll.blog.com/internal/controller"
	middleware "zll.blog.com/internal/middleWare"
	"zll.blog.com/internal/pb"
)

func RegisterUserFollowRouter(r *gin.Engine, conn *grpc.ClientConn) {
	userFollowGroup := r.Group("/userFollow")

	userFollowClient := pb.NewUserFollowServiceClient(conn)
	userFollowController := controller.NewUserFollowController(userFollowClient)
	{
		userFollowGroup.GET("/follower/:followingId", userFollowController.GetUserFollowerIds)
		userFollowGroup.GET("/following/:followerId", userFollowController.GetUserFollowingIds)
	}
	userFollowAuthGroup := userFollowGroup.Group("/auth")
	userFollowAuthGroup.Use(middleware.JWTAuthMiddleWare(config.GetConfig().Jwt.SecretKey))
	{
		userFollowAuthGroup.GET("/isFollow/:followerId/:followingId", userFollowController.GetIsFollow)
		userFollowAuthGroup.PUT("/isFollow/:followerId/:followingId/:isFollow", userFollowController.SetIsFollow)
	}
}
