package model

type User struct {
	ID       int    `gorm:"type:int;primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	UserName string `gorm:"type:varchar(20)" json:"username"`
	Age      int    `gorm:"type:int" json:"age"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Phone    string `gorm:"type:varchar(15)" json:"phone"`
	Tags     []Tag  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tags"`
}
