package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/config"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
	"zll.blog.com/internal/util"
)

type userController struct {
	userClient pb.UserServiceClient
}

func NewUserController(userClient pb.UserServiceClient) *userController {
	return &userController{userClient: userClient}
}

func (uc *userController) GetUserDTOById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := uc.userClient.GetUserDTOById(context.Background(), &pb.UserIdRequest{Id: uint64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch userDTO by id"})
		log.Fatalf("failed to get a userDTO by id: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": converter.ProtoToModelForUserDTO(res)})
}

func (uc *userController) Login(c *gin.Context) {
	var userLoginRequest *pb.UserLoginRequest
	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Invalid JSON"})
		return
	}

	res, err := uc.userClient.Login(context.Background(), userLoginRequest)
	if err != nil || res.Id == 0 {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": "密码错误"})
		return
	}

	user := converter.ProtoToModelForUser(res)
	if !user.DeletedAt.Time.IsZero() {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": "该账户已注销"})
		return
	}

	config := config.GetConfig()
	token, err := util.GenerateToken(uint(res.Id), config.Jwt.SecretKey, config.Jwt.Signed, config.Jwt.ExpireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "未知错误"})
		return
	}

	redis := util.GetRedis(config.Redis.Addr, config.Redis.Password, config.Redis.Db)
	if err := redis.Set(strconv.Itoa(int(user.Id)), token, config.Jwt.ExpireTime*time.Hour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "未知错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func (uc *userController) GetUserById(c *gin.Context) {
	value, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt已过期"})
	}
	id, exists := value.(uint)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt已过期"})
	}
	res, err := uc.userClient.GetUserById(context.Background(), &pb.UserIdRequest{Id: uint64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt已过期"})
	}
	c.JSON(http.StatusOK, gin.H{"data": converter.ProtoToModelForUser(res)})
}

func (uc *userController) GetUserDTOsByKeywords(c *gin.Context) {
	keywords := c.Param("keywords")
	userDTOs, err := uc.userClient.GetUserDTOsByKeywords(context.Background(), &pb.UserKeywordsRequest{Keywords: keywords})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
	}
	var uDs []*model.UserDTO
	for _, u := range userDTOs.UserDTOs {
		uDs = append(uDs, converter.ProtoToModelForUserDTO(u))
	}
	c.JSON(http.StatusOK, gin.H{"data": uDs})
}

func (uc *userController) Register(c *gin.Context) {
	var userLoginRequest *pb.UserLoginRequest
	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Invalid JSON"})
		return
	}

	res, err := uc.userClient.Register(context.Background(), userLoginRequest)
	if err != nil || res.Id == 0 {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": "注册失败"})
		return
	}

	res, err = uc.userClient.Login(context.Background(), userLoginRequest)
	if err != nil || res.Id == 0 {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": "密码错误"})
		return
	}

	user := converter.ProtoToModelForUser(res)
	if !user.DeletedAt.Time.IsZero() {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": "该账户已注销"})
		return
	}

	config := config.GetConfig()
	token, err := util.GenerateToken(uint(res.Id), config.Jwt.SecretKey, config.Jwt.Signed, config.Jwt.ExpireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "未知错误"})
		return
	}

	redis := util.GetRedis(config.Redis.Addr, config.Redis.Password, config.Redis.Db)
	if err := redis.Set(strconv.Itoa(int(user.Id)), token, config.Jwt.ExpireTime*time.Hour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "未知错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}
