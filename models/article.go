package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title     string
	Content   string
	Summary   string
	Thumbnail string
	IsTop     bool
	Status    bool
	ViewCount int64
	IsComment bool
}