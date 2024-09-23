package handlers

import (
	"blog/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 创建标签
// @Produce json
// @Success 200
// @Router /tag [post]
func CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.Bind(&tag); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg: "获取创建标签信息失败",
		})
		return
	}

	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()

	if result := db.Create(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg: "创建标签失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg: "创建标签成功",
	})
}