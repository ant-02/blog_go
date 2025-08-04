package model

import (
	"time"
)

type Article struct {
	BaseModel
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	CoverImage string `json:"coverImage"`
	AuthorId   uint64 `json:"authorId"`
	Status     string `json:"status"`
	Content    string `json:"content"`
}

type ArticleDTO struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	CoverImage string    `json:"coverImage"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Summary    string    `json:"summary"`
}

type ArticleDTOs struct {
	ArticleDTOs []*ArticleDTO
	Count       int64
}
