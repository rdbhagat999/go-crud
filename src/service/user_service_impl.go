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
func (u *UserServiceImpl) Create(user request.CreateUserRequest) response.UserResponse {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)

	userModel := model.User{
		Name:     user.Name,
		Username: user.Username,
		Age:      user.Age,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	result, resultErr := u.UserRepository.Save(userModel)
	helper.ErrorPanic(resultErr)

	userReponse := response.UserResponse{
		ID:       result.ID,
		Name:     result.Name,
		Username: result.Username,
		Age:      result.Age,
		Email:    result.Email,
		Phone:    result.Phone,
		Tags:     result.Tags,
		Posts:    result.Posts,
	}

	return userReponse
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(UserId int) {
	u.UserRepository.Delete(UserId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() []response.UserResponse {
	var users = []response.UserResponse{}
	result, err := u.UserRepository.FindAll()
	helper.ErrorPanic(err)

	for _, v := range result {
		found := response.UserResponse{
			ID:       v.ID,
			Name:     v.Name,
			Username: v.Username,
			Age:      v.Age,
			Email:    v.Email,
			Phone:    v.Phone,
			Tags:     v.Tags,
			Posts:    v.Posts,
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
		Username: result.Username,
		Age:      result.Age,
		Email:    result.Email,
		Phone:    result.Phone,
		Tags:     result.Tags,
		Posts:    result.Posts,
	}

	return userReponse
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) response.UserResponse {
	found, err := u.UserRepository.FindById(user.ID)
	helper.ErrorPanic(err)

	found.Name = user.Name
	// found.Username = user.Username
	found.Age = user.Age
	found.Email = user.Email
	found.Phone = user.Phone

	result, resultErr := u.UserRepository.Update(found)
	helper.ErrorPanic(resultErr)

	userReponse := response.UserResponse{
		ID:       result.ID,
		Name:     result.Name,
		Username: result.Username,
		Age:      result.Age,
		Email:    result.Email,
		Phone:    result.Phone,
		Tags:     result.Tags,
		Posts:    result.Posts,
	}

	return userReponse

}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
