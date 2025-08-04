package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/config"
	"zll.blog.com/internal/util"
)

func JWTAuthMiddleWare(secretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "jwt已过期"})
			ctx.Abort()
			return
		}

		token, err := util.ParseToken(tokenString, secretKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "无效的认证令牌"})
			ctx.Abort()
			return
		}

		config := config.GetConfig()
		redis := util.GetRedis(config.Redis.Addr, config.Redis.Password, config.Redis.Db)
		val, err := redis.Get(strconv.Itoa(int(token.UserId)))
		if err != nil || val != tokenString {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "无效的认证令牌"})
			ctx.Abort()
			return
		}

		ctx.Set("userId", token.UserId)
		ctx.Next()
	}
}
