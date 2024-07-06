package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID        int            `gorm:"type:int UNSIGNED NOT NULL AUTO_INCREMENT;primarKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);uniqueIndex" json:"name"`
	UserID    int            `gorm:"type:int" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
