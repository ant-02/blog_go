package handlers

import (
	"blog/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 获取文章预览信息
// @Produce json
// @Success 200
// @Router /articleTags [get]
func GetArticlePreviewTags(c *gin.Context) {
	var tags []models.Tag
	if result := db.Find(&tags); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg: "获取文章预览失败",
		})
		return
	}

	var articlePreviewTagsRes models.ArticlePreviewTagsRes
	articlePreviewTagsRes.ArticlePreviewTags = make([]models.ArticlePreviewTag, len(tags))
	articlePreviewTagsRes.Size = len(tags)
	for i, tag := range tags {
		var articleTags []models.ArticleTag
		if result := db.Where("tag_id", tag.ID).Find(&articleTags); result.Error != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Code: 500,
				Msg: "获取文章预览失败",
			})
			return
		}

		articlePreviewTagsRes.ArticlePreviewTags[i].ID = tag.ID
		articlePreviewTagsRes.ArticlePreviewTags[i].Name = tag.Name
		articlePreviewTagsRes.ArticlePreviewTags[i].ArticlePreviews = make([]models.ArticlePreview, len(articleTags))
		for j, articleTag := range articleTags {
			var article models.Article
			if result := db.Order("is_deleted desc").Order("created_at desc").Limit(6).Find(&article, articleTag.ArticleId); result.Error != nil {
				c.JSON(http.StatusInternalServerError, models.Response{
					Code: 500,
					Msg: "获取文章预览失败",
				})
				return
			}
			articlePreviewTagsRes.ArticlePreviewTags[i].ArticlePreviews[j] = article.GenerateArticlePreview()
		}
	}
	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg: "获取文章预览成功",
		Data: articlePreviewTagsRes,
	})
}

func GetArticlePreviewTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	if result := db.First(&tag, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg: "获取标签文章失败",
		})
		return 
	}

	var articleTags []models.ArticleTag
	if result := db.Where("tag_id", tag.ID).Find(&articleTags); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code: 500,
			Msg: "获取标签文章失败",
		})
		return
	}

	var articlePreviewTag models.ArticlePreviewTag
	articlePreviewTag.ID = tag.ID
	articlePreviewTag.Name = tag.Name
	articlePreviewTag.ArticlePreviews = make([]models.ArticlePreview, len(articleTags))
	for i, articleTag := range articleTags {
		fmt.Println(articleTag)
		var article models.Article
		if result := db.First(&article, articleTag.ArticleId); result.Error != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Code: 500,
				Msg: "获取标签文章失败",
			})
			return
		}
		articlePreviewTag.ArticlePreviews[i] = article.GenerateArticlePreview()
	}

	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg: "获取标签文章成功",
		Data: articlePreviewTag,
	})
}