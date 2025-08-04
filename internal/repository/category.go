package repository

import (
	"gorm.io/gorm"
	"zll.blog.com/internal/model"
)

type CategoryRepository interface {
	GetAll() ([]*model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db gorm.DB) CategoryRepository {
	return &categoryRepository{db: &db}
}

func (cr *categoryRepository) GetAll() ([]*model.Category, error) {
	var categories []*model.Category
	err := cr.db.Where("deleted_at IS NULL").
		Find(&categories).Error
	return categories, err
}
