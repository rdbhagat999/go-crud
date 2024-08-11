package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	// gorm.Model
	ID        int            `gorm:"type:int UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(30);uniqueIndex" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
