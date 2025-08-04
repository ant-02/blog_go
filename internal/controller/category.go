package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/pb"
)

type categoryController struct {
	categoryClient pb.CategoryServiceClient
}

func NewCategoryController(categoryClient pb.CategoryServiceClient) *categoryController {
	return &categoryController{categoryClient: categoryClient}
}

func (cc *categoryController) GetAll(c *gin.Context) {
	res, err := cc.categoryClient.GetAll(context.Background(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch categories"})
		log.Fatalf("failed to get all categories: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": converter.ProtoToModelForCategoryList(res)})
}
