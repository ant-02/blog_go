package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/pb"
)

type userFollowController struct {
	userFollowClient pb.UserFollowServiceClient
}

func NewUserFollowController(userFollowClient pb.UserFollowServiceClient) *userFollowController {
	return &userFollowController{userFollowClient: userFollowClient}
}

func (ufc *userFollowController) GetUserFollowerIds(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("followingId"))
	ids, err := ufc.userFollowClient.GetUserFollowerIds(context.Background(), &pb.UserFollowIdRequest{Id: uint64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ids.Id})
}

func (ufc *userFollowController) GetUserFollowingIds(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("followerId"))
	ids, err := ufc.userFollowClient.GetUserFollowingIds(context.Background(), &pb.UserFollowIdRequest{Id: uint64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ids.Id})
}

func (ufc *userFollowController) GetIsFollow(c *gin.Context) {
	followerId, _ := strconv.Atoi(c.Param("followerId"))
	followingId, _ := strconv.Atoi(c.Param("followingId"))
	value, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt已过期"})
	}
	id, exists := value.(uint)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt已过期"})
	}
	if id != uint(followerId) {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据错误"})
	}
	res, err := ufc.userFollowClient.GetIsFollow(context.Background(), &pb.IsFollowedRequest{FollowerId: uint64(followerId), FollowingId: uint64(followingId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res.IsFollow})
}

func (ufc *userFollowController) SetIsFollow(c *gin.Context) {
	followerId, _ := strconv.Atoi(c.Param("followerId"))
	followingId, _ := strconv.Atoi(c.Param("followingId"))
	isFollow, _ := strconv.Atoi(c.Param("isFollow"))
	value, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt已过期"})
	}
	id, exists := value.(uint)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt已过期"})
	}
	if id != uint(followerId) {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据错误"})
	}
	res, err := ufc.userFollowClient.SetIsFollow(context.Background(), &pb.IsFollowedRequest{FollowerId: uint64(followerId), FollowingId: uint64(followingId), IsFollow: isFollow != 0})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res.IsFollow})
}
