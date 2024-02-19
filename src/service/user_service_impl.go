package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/model"
	"go-crud/src/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)

	userModel := model.User{
		Name:     user.Name,
		UserName: user.UserName,
		Age:      user.Age,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	u.UserRepository.Save(userModel)
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(UserId int) {
	u.UserRepository.Delete(UserId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() []response.UserResponse {
	var users []response.UserResponse
	result, err := u.UserRepository.FindAll()
	helper.ErrorPanic(err)

	for _, v := range result {
		found := response.UserResponse{
			ID:       v.ID,
			Name:     v.Name,
			UserName: v.UserName,
			Age:      v.Age,
			Email:    v.Email,
			Phone:    v.Phone,
			Tags:     v.Tags,
		}

		users = append(users, found)
	}

	return users
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	result, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	userReponse := response.UserResponse{
		ID:       result.ID,
		Name:     result.Name,
		UserName: result.UserName,
		Age:      result.Age,
		Email:    result.Email,
		Phone:    result.Phone,
		Tags:     result.Tags,
	}

	return userReponse
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	found, err := u.UserRepository.FindById(user.ID)
	helper.ErrorPanic(err)

	found.Name = user.Name
	// found.UserName = user.UserName
	found.Age = user.Age
	found.Email = user.Email
	found.Phone = user.Phone

	u.UserRepository.Update(found)
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
