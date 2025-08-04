package repository

import (
	"gorm.io/gorm"
	"zll.blog.com/internal/model"
)

type articleTagRepository struct {
	db *gorm.DB
}

type ArticleTagRepository interface {
	GetIdsByArticleId(id uint) ([]uint, error)
}

func NewArticleTagRepository(db *gorm.DB) ArticleTagRepository {
	return &articleTagRepository{db: db}
}

func (tr *articleTagRepository) GetIdsByArticleId(id uint) ([]uint, error) {
	var ids []uint
	err := tr.db.Model(&model.ArticleTag{}).
		Select("tag_id").
		Where("article_id = ?", id).
		Find(&ids).Error
	return ids, err
}
