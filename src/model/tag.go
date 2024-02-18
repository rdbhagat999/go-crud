package model

type Tag struct {
	Id     int    `gorm:"type:int;primarKey"`
	Name   string `gorm:"type:varchar(255)"`
	UserID int    `gorm:"type:int"`
}
