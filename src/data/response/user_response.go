package response

import "go-crud/src/model"

type UserResponse struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Age      int          `json:"age"`
	RoleID   int          `json:"role_id"`
	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	Posts    []model.Post `json:"posts"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
