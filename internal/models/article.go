package models

import "gorm.io/gorm"

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
