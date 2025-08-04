package service

import (
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/repository"
)

type ArticleService interface {
	GetArticleById(id uint) (model.Article, error)
	GetArticleDTOById(id uint) (*model.ArticleDTO, error)
	GetArticleDTOsByKeywords(keywords string) ([]*model.ArticleDTO, error)
	GetArticleDTOsByUserId(userId uint, page, pageSize int, status string) (*model.ArticleDTOs, error)
	SaveArticle(article *model.Article) (bool, error)
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(articleRepository repository.ArticleRepository) ArticleService {
	return &articleService{articleRepository: articleRepository}
}

func (s *articleService) GetArticleById(id uint) (model.Article, error) {
	return s.articleRepository.GetById(id)
}

func (s *articleService) GetArticleDTOById(id uint) (*model.ArticleDTO, error) {
	return s.articleRepository.GetArticleDTOById(id)
}

func (s *articleService) GetArticleDTOsByKeywords(keywords string) ([]*model.ArticleDTO, error) {
	return s.articleRepository.GetArticleDTOsByKeywords(keywords)
}

func (s *articleService) GetArticleDTOsByUserId(userId uint, page, pageSize int, status string) (*model.ArticleDTOs, error) {
	return s.articleRepository.GetArticleDTOsByUserId(userId, page, pageSize, status)
}

func (s *articleService) SaveArticle(article *model.Article) (bool, error) {
	return s.articleRepository.SaveArticle(article)
}
