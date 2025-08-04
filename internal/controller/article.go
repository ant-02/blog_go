package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
)

type articleController struct {
	articleClient         pb.ArticleServiceClient
	articleCategoryClient pb.ArticleCategoryServiceClient
}

func NewArticleController(articleClient pb.ArticleServiceClient, articleCategoryClient pb.ArticleCategoryServiceClient) *articleController {
	return &articleController{articleClient: articleClient, articleCategoryClient: articleCategoryClient}
}

// @Summary 通过id获取文章
// @Produce json
// @Success 200
// @Param id path int true "用户 ID"
// @Router /article [get]
func (ac *articleController) GetArticleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ac.articleClient.GetArticleById(context.Background(), &pb.ArticleIdRequest{Id: uint64(id)})
	if res.Article.Id == 0 || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch article"})
		log.Fatalf("failed to get a article by id: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": converter.ProtoToModelForArticle(res.Article)})
}

func (ac *articleController) GetArticleDTOsByCategoryId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	count, _ := strconv.Atoi(c.Param("count"))

	articleCategoryList, err := ac.articleCategoryClient.GetByCategoryId(context.Background(), &pb.ArticleCategoryIdRequest{Id: uint64(id), Count: int64(count)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		log.Fatalf("failed to get articleCategoryList by categoryId: %v", err)
	}
	acl := converter.ProtoToModelForArticleCategoryList(articleCategoryList)
	var articleDTOs []*model.ArticleDTO
	for _, articleCategory := range acl {
		articleDTO, err := ac.articleClient.GetArticleDTOById(context.Background(), &pb.ArticleIdRequest{Id: uint64(articleCategory.ArticleId)})
		if articleDTO.Id == 0 {
			continue
		}
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			log.Fatalf("failed to get articleCategoryList by categoryId: %v", err)
		}
		articleDTOs = append(articleDTOs, converter.ProtoToModelForArticleDTO(articleDTO))
	}
	c.JSON(http.StatusOK, gin.H{"data": articleDTOs})
}

func (ac *articleController) GetArticleDTOsByKeywords(c *gin.Context) {
	keywords := c.Param("keywords")
	articleDTOs, err := ac.articleClient.GetArticleDTOsByKeyWords(context.Background(), &pb.ArticleKeywordsRequest{Keywords: keywords})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": err})
		return
	}
	var aDs []*model.ArticleDTO
	for _, a := range articleDTOs.ArticleDTOs {
		aDs = append(aDs, converter.ProtoToModelForArticleDTO(a))
	}
	c.JSON(http.StatusOK, gin.H{"data": aDs})
}

func (ac *articleController) GetArticleDTOsByUserId(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	page, _ := strconv.Atoi(c.Param("page"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	status := c.Param("status")
	articleDTOs, err := ac.articleClient.GetArticleDTOsByUserId(context.Background(), &pb.ArticleIdRequest{Id: uint64(userId), Page: uint64(page), PageSize: uint64(pageSize), Status: status})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": err})
		return
	}
	var aDs []*model.ArticleDTO
	for _, a := range articleDTOs.ArticleDTOs {
		aDs = append(aDs, converter.ProtoToModelForArticleDTO(a))
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"articleDTOs": aDs, "count": articleDTOs.Count}})
}

func (ac *articleController) SaveArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "0"})
		return
	}
	res, err := ac.articleClient.SaveArticle(context.Background(), &pb.ArticleRequest{Article: converter.ModelToProtoForArticle(&article)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "0"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
