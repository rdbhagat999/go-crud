package request

type CreateUserRequest struct {
	Name     string `validate:"required,min=5,max=200" json:"name"`
	Username string `validate:"required,min=5,max=200" json:"username"`
	Password string `validate:"required,min=5,max=255" json:"password"`
	Age      int    `validate:"required,min=18,max=60" json:"age"`
	Email    string `validate:"required,min=5,max=200" json:"email"`
	Phone    string `validate:"required,min=10,max=15" json:"phone"`
}

type UpdateUserRequest struct {
	ID    int    `validate:"required" json:"user_id"`
	Name  string `validate:"required,min=5,max=200" json:"name"`
	Age   int    `validate:"required,min=18,max=60" json:"age"`
	Email string `validate:"required,min=5,max=200" json:"email"`
	Phone string `validate:"required,min=10,max=15" json:"phone"`
}

type LoginUserRequest struct {
	Username string `validate:"required,min=5,max=200" json:"username"`
	Password string `validate:"required,min=5,max=255" json:"password"`
}
