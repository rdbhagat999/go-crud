package request

type UpdateUserRequest struct {
	Id   int    `validate:"required" json:"user_id"`
	Name string `validate:"required,min=5,max=200" json:"name"`
	// UserName string `validate:"required,min=5,max=200" json:"username"`
	Age   int    `validate:"required,min=18,max=60" json:"age"`
	Email string `validate:"required,min=5,max=200" json:"email"`
	Phone string `validate:"required,min=10,max=15" json:"phone"`
}
