package service

import (
	"zll.blog.com/internal/repository"
)

type articleTagService struct {
	articleTagRepository repository.ArticleTagRepository
}

type ArticleTagService interface {
	GetIdsByArticleId(id uint) ([]uint, error)
}

func NewArticleTagService(articleTagRepository repository.ArticleTagRepository) ArticleTagService {
	return &articleTagService{articleTagRepository: articleTagRepository}
}

func (ts *articleTagService) GetIdsByArticleId(id uint) ([]uint, error) {
	return ts.articleTagRepository.GetIdsByArticleId(id)
}
