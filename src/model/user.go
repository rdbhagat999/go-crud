package model

type User struct {
	ID       int    `gorm:"type:int;primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"type:varchar(20);uniqueIndex" json:"username"`
	Age      int    `gorm:"type:int" json:"age"`
	Email    string `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Phone    string `gorm:"type:varchar(255);uniqueIndex" json:"phone"`
	Tags     []Tag  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tags"`
}
