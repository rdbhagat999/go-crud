package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
)

type UserService interface {
	Login(user request.LoginUserRequest) (userdata response.UserResponse, err error)
	FindByUsername(username string) (userdata response.UserResponse, err error)
	Create(user request.CreateUserRequest) (userdata response.UserResponse, err error)
	Update(id int, user request.UpdateUserRequest) (userdata response.UserResponse, err error)
	Delete(userId int)
	FindById(userId int) (userdata response.UserResponse, err error)
	FindAll() (userdata []response.UserResponse, err error)
}
