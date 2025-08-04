package service

import (
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/repository"
)

type CategoryService interface {
	GetAll() ([]*model.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (cs *categoryService) GetAll() ([]*model.Category, error) {
	return cs.categoryRepository.GetAll()
}
