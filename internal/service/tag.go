package service

import (
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/repository"
)

type tagService struct {
	tagRepository repository.TagRepository
}

type TagService interface {
	GetAll() ([]*model.Tag, error)
}

func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{tagRepository: tagRepository}
}

func (ts *tagService) GetAll() ([]*model.Tag, error) {
	return ts.tagRepository.GetAll()
}
