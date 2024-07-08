package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	// gorm.Model
	ID        uint           `gorm:"type:int UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(255);uniqueIndex" json:"title"`
	Body      string         `gorm:"type:varchar(255);" json:"body"`
	UserID    uint           `gorm:"type:int" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
