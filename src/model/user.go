package model

type User struct {
	Id       int    `gorm:"type:int;primaryKey"`
	Name     string `gorm:"type:varchar(255)"`
	UserName string `gorm:"type:varchar(20)"`
	Age      int    `gorm:"type:int"`
	Email    string `gorm:"type:varchar(255)"`
	Phone    string `gorm:"type:varchar(15)"`
	Tags     []Tag  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
