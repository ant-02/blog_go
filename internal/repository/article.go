package repository

import (
	"time"

	"gorm.io/gorm"
	"zll.blog.com/internal/model"
)

type ArticleRepository interface {
	GetById(id uint) (model.Article, error)
	GetArticleDTOById(id uint) (*model.ArticleDTO, error)
	GetArticleDTOsByKeywords(keywords string) ([]*model.ArticleDTO, error)
	GetArticleDTOsByUserId(userId uint, page, pageSize int, status string) (*model.ArticleDTOs, error)
	SaveArticle(article *model.Article) (bool, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) GetById(id uint) (model.Article, error) {
	var article model.Article
	err := r.db.Find(&article, id).Error
	return article, err
}

func (r *articleRepository) GetArticleDTOById(id uint) (*model.ArticleDTO, error) {
	var articleDTO *model.ArticleDTO
	err := r.db.Model(&model.Article{}).
		Select("id", "title", "updated_at", "cover_image").
		Where("status = 'published'").
		Find(&articleDTO, id).Error
	return articleDTO, err
}

func (r *articleRepository) GetArticleDTOsByKeywords(keywords string) ([]*model.ArticleDTO, error) {
	var articleDTOs []*model.ArticleDTO
	keywords = "%" + keywords + "%"
	err := r.db.Model(&model.Article{}).
		Select("id", "title", "updated_at", "cover_image", "summary").
		Where("status = 'published'").
		Where("title LIKE ?", keywords).
		Find(&articleDTOs).Error
	return articleDTOs, err
}

func (r *articleRepository) GetArticleDTOsByUserId(userId uint, page, pageSize int, status string) (*model.ArticleDTOs, error) {
	var res *model.ArticleDTOs = &model.ArticleDTOs{}
	offset := (page - 1) * pageSize
	var where string
	if status == "all" {
		where = "author_id = ?"
	} else if status == "published" {
		where = "status = 'published' AND author_id = ?"
	}
	err := r.db.Model(&model.Article{}).
		Select("id", "title", "updated_at", "cover_image", "summary").
		Where(where, userId).
		Count(&res.Count).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&model.Article{}).
		Select("id", "title", "updated_at", "cover_image", "summary").
		Where(where, userId).
		Offset(offset).
		Limit(pageSize).
		Find(&res.ArticleDTOs).Error
	return res, err
}

func (r *articleRepository) SaveArticle(article *model.Article) (bool, error) {
	article.UpdatedAt = time.Now()
	var err error
	if article.Id == 0 {
		article.CreatedAt = time.Now()
		err = r.db.
			Omit("deleted_at").
			Save(article).Error
	} else {
		err = r.db.
			Omit("created_at", "deleted_at").
			Save(article).Error
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
