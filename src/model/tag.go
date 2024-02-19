package model

type Tag struct {
	ID     int    `gorm:"type:int;primarKey" json:"id"`
	Name   string `gorm:"type:varchar(255)" json:"name"`
	UserID int    `gorm:"type:int" json:"user_id"`
}
