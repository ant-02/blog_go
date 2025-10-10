package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uint64    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
