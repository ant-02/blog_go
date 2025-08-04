package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/pb"
)

type tagController struct {
	tagClient pb.TagServiceClient
}

func NewTagController(tagClient pb.TagServiceClient) *tagController {
	return &tagController{tagClient: tagClient}
}

func (tc *tagController) GetAll(c *gin.Context) {
	res, err := tc.tagClient.GetAll(context.Background(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch tags"})
		log.Fatalf("failed to get all tags: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": converter.ProtoToModelForTagList(res)})
}
