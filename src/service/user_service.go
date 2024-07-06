package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest) response.UserResponse
	Update(user request.UpdateUserRequest) response.UserResponse
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
}
