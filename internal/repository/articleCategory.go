package repository

import (
	"gorm.io/gorm"
	"zll.blog.com/internal/model"
)

type ArticleCategoryRepository interface {
	GetByCategoryId(id uint, count int) ([]*model.ArticleCategory, error)
	GetCategoryIdByArticleId(id uint) (uint, error)
	SaveArticleCategory(articleCategory *model.ArticleCategory) (bool, error)
}

type articleCategoryRepository struct {
	db *gorm.DB
}

func NewArticleCategoryRepository(db *gorm.DB) ArticleCategoryRepository {
	return &articleCategoryRepository{db: db}
}

func (acr *articleCategoryRepository) GetByCategoryId(id uint, count int) ([]*model.ArticleCategory, error) {
	var articleCategoryList []*model.ArticleCategory
	err := acr.db.Where("category_id = ?", id).
		Limit(count).
		Find(&articleCategoryList).Error
	return articleCategoryList, err
}

func (acr *articleCategoryRepository) GetCategoryIdByArticleId(id uint) (uint, error) {
	var res uint
	err := acr.db.Model(&model.ArticleCategory{}).
		Select("category_id").
		Where("article_id = ?", id).
		Find(&res).Error
	return res, err
}

func (c *articleCategoryRepository) SaveArticleCategory(articleCategory *model.ArticleCategory) (bool, error) {
	id, _ := c.GetCategoryIdByArticleId(articleCategory.ArticleId)
	var err error = nil
	if id == 0 {
		err = c.db.Model(&model.ArticleCategory{}).
			Create(articleCategory).Error
		if err != nil {
			return false, err
		}
	} else if id != articleCategory.CategoryId {
		err = c.db.
			Model(&model.ArticleCategory{}).
			Where("article_id = ?", articleCategory.ArticleId).
			Update("category_id", articleCategory.CategoryId).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
