package models

import "gorm.io/gorm"

var Tag struct {
	gorm.Model
	Name string
}