package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint           `gorm:"type:int UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	Username  string         `gorm:"type:varchar(20);uniqueIndex" json:"username"`
	Password  []byte         `gorm:"type:varchar(255);" json:"password"`
	Age       uint           `gorm:"type:int" json:"age"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Phone     string         `gorm:"type:varchar(255);uniqueIndex" json:"phone"`
	Tags      []Tag          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tags"`
	Posts     []Post         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"posts"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
