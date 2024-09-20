package handlers

import (
	"blog/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	var articles []models.Article
	result := db.Find(&articles)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg:  "获取全部文章信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg:  "获取全部文章信息成功",
		Data: articles,
	})
}

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
