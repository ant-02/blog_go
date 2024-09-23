package handlers

import (
	"blog/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 创建新的文章
// @Produce json
// @Success 200
// @Router /article [post]
func CreateArticle(c *gin.Context) {
	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg:  "接收文章信息有误",
		})
		return
	}

	article.ViewCount = 0
	article.IsDeleted = false
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	result := db.Create(&article)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg:  "创建文章信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg:  "获取全部文章信息成功",
	})
}

// @Summary 通过id获取文章
// @Produce json
// @Success 200
// @Router /article [get]
func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	result := db.First(&article, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg:  "通过id查询文章信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg:  "通过id查询文章信息成功",
		Data: article,
	})
}

// @Summary 获取文章预览信息
// @Produce json
// @Success 200
// @Router /articles [get]
func GetArticlePreview(c *gin.Context) {
	var articles []models.Article
	
	result := db.Order("is_top desc").Order("created_at desc").Find(&articles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg: "获取文章预览失败",
		})
		return
	}
	

	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg: "获取文章预览成功",
		Data: models.MakeArticlePreviewRes(articles),
	})
}