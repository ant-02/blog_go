package service

import (
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/repository"
)

type ArticleCategoryService interface {
	GetByCategoryId(id uint, count int) ([]*model.ArticleCategory, error)
	GetCategoryIdByArticleId(id uint) (uint, error)
	SaveArticleCategory(articleCategory *model.ArticleCategory) (bool, error)
}

type articleCategoryService struct {
	articleCategoryRepo repository.ArticleCategoryRepository
}

func NewArticleCategoryService(articleCategoryRepo repository.ArticleCategoryRepository) *articleCategoryService {
	return &articleCategoryService{articleCategoryRepo: articleCategoryRepo}
}

func (acs *articleCategoryService) GetByCategoryId(id uint, count int) ([]*model.ArticleCategory, error) {
	return acs.articleCategoryRepo.GetByCategoryId(id, count)
}

func (acs *articleCategoryService) GetCategoryIdByArticleId(id uint) (uint, error) {
	return acs.articleCategoryRepo.GetCategoryIdByArticleId(id)
}

func (acs *articleCategoryService) SaveArticleCategory(articleCategory *model.ArticleCategory) (bool, error) {
	return acs.articleCategoryRepo.SaveArticleCategory(articleCategory)
}