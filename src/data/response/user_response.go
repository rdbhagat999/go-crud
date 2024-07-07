package response

import "go-crud/src/model"

type UserResponse struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Age      int          `json:"age"`
	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	Tags     []model.Tag  `json:"tags"`
	Posts    []model.Post `json:"posts"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
