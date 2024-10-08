package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        int            `gorm:"type:int UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	Username  string         `gorm:"type:varchar(20);unique" json:"username"`
	Password  []byte         `gorm:"type:varchar(255);" json:"password"`
	Age       int            `gorm:"type:int" json:"age"`
	Email     string         `gorm:"type:varchar(255);unique" json:"email"`
	Phone     string         `gorm:"type:varchar(255);unique" json:"phone"`
	RoleID    *int           `gorm:"type:int;default:1" json:"role_id"`
	Posts     []Post         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"posts"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
