package repository

import (
	"gorm.io/gorm"
	"zll.blog.com/internal/model"
)

type tagRepository struct {
	db *gorm.DB
}

type TagRepository interface {
	GetAll() ([]*model.Tag, error)
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (tr *tagRepository) GetAll() ([]*model.Tag, error) {
	var tags []*model.Tag
	err := tr.db.Where("deleted_at IS NULL").Find(&tags).Error
	return tags, err
}
