package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
)

type articleCategoryController struct {
	articleCategoryClient pb.ArticleCategoryServiceClient
}

func NewArticleCategoryController(articleCategoryClient pb.ArticleCategoryServiceClient) *articleCategoryController {
	return &articleCategoryController{articleCategoryClient: articleCategoryClient}
}

func (acc *articleCategoryController) GetCategoryIdByArticleId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := acc.articleCategoryClient.GetCategoryIdByArticleId(context.Background(), &pb.ArticleCategoryIdRequest{Id: uint64(id)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res.Id})
}

func (acc *articleCategoryController) SaveArticleCategory(c *gin.Context) {
	var articleCategory model.ArticleCategory
	if err := c.ShouldBindJSON(&articleCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "0"})
		return
	}
	res, err := acc.articleCategoryClient.SaveArticleCategory(context.Background(), &pb.ArticleCategoryRequest{ArticleCategory: converter.ModelToProtoForArticleCategory(&articleCategory)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res.Ok})
}
