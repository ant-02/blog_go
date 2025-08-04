package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/pb"
)

type articleTagController struct {
	articleTagClient pb.ArticleTagServiceClient
}

func NewArticleTagController(articleTagClient pb.ArticleTagServiceClient) *articleTagController {
	return &articleTagController{articleTagClient: articleTagClient}
}

func (tc *articleTagController) GetTagIdsByArticleId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := tc.articleTagClient.GetTagIdsByArticleId(context.Background(), &pb.ArticleTagIdRequest{Id: uint64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch tagIds by articleId"})
		log.Fatalf("failed to get tagIds by articleId: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": res.Ids})
}
