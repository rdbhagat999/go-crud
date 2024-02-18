package response

import "go-crud/src/model"

type UserResponse struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	UserName string      `json:"username"`
	Age      int         `json:"age"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	Tags     []model.Tag `json:"tags"`
}
