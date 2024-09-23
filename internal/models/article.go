package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title     string `json:"title"`
	Content   string `json:"content"`
	Summary   string `json:"summary"`
	Thumbnail string `json:"thumbnail"`
	IsTop     bool   `json:"isTop"`
	Status    bool   `json:"status"`
	ViewCount int64  `json:"viewCount"`
	IsComment bool   `json:"isComment"`
	IsDeleted bool   `json:"isDeleted"`
}

type ArticlePreview struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Thumbnail string    `json:"thumbnail"`
	IsTop     bool      `json:"isTop"`
	ViewCount int64     `json:"viewCount"`
	CreatedAt time.Time `json:"createdAt"`
}



func (a *Article) GenerateArticlePreview() ArticlePreview {
	return ArticlePreview{
		ID:        a.ID,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		IsTop:     a.IsTop,
		ViewCount: a.ViewCount,
		CreatedAt: a.CreatedAt,
	}
}
