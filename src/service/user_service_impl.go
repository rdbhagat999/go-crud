package service

import (
	"fmt"
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/model"
	"go-crud/src/repository"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func userServicePrintln(err error) {
	fmt.Println("userServicePrintln: " + err.Error())
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) (userdata response.UserResponse, err error) {
	var userReponse response.UserResponse

	validateErr := u.Validate.Struct(user)
	// helper.ErrorPanic(err)

	if validateErr != nil {
		userServicePrintln(validateErr)
		//  validateErr.(validator.ValidationErrors)
		return userReponse, validateErr
	}

	password, passErr := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	// helper.ErrorPanic(passErr)

	if passErr != nil {
		userServicePrintln(passErr)
		return userReponse, passErr
	}

	userModel := model.User{
		Name:     user.Name,
		Username: user.Username,
		Age:      user.Age,
		RoleID:   &user.RoleID,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: password,
	}

	result, resultErr := u.UserRepository.Save(userModel)
	// helper.ErrorPanic(resultErr)

	if resultErr != nil {
		userServicePrintln(resultErr)
		return userReponse, resultErr
	}

	userReponse = response.UserResponse{
		ID:       int(result.ID),
		Name:     result.Name,
		Username: result.Username,
		Age:      int(result.Age),
		RoleID:   int(*result.RoleID),
		Email:    result.Email,
		Phone:    result.Phone,
		Posts:    result.Posts,
	}

	return userReponse, nil
}

// Login implements UserService.
func (u *UserServiceImpl) Login(user request.LoginUserRequest) (userdata response.UserResponse, err error) {
	var userReponse response.UserResponse

	validateErr := u.Validate.Struct(user)
	// helper.ErrorPanic(validateErr)

	if validateErr != nil {
		userServicePrintln(validateErr)
		return userReponse, validateErr
	}

	result, loginErr := u.UserRepository.Login(user)
	// helper.ErrorPanic(loginErr)

	if loginErr != nil {
		userServicePrintln(loginErr)
		return userReponse, loginErr
	}

	compareErr := bcrypt.CompareHashAndPassword(result.Password, []byte(user.Password))
	// helper.ErrorPanic(compareErr)

	if compareErr != nil {
		userServicePrintln(compareErr)
		return userReponse, compareErr
	}

	// loadConfig, loadConfigErr := config.LoadConfig("../../app.env")
	// helper.ErrorPanic(loadConfigErr)

	userReponse = response.UserResponse{
		ID:       int(result.ID),
		Name:     result.Name,
		Username: result.Username,
		Age:      int(result.Age),
		RoleID:   int(*result.RoleID),
		Email:    result.Email,
		Phone:    result.Phone,
		Posts:    result.Posts,
	}

	return userReponse, nil
}

// FindByUsername implements UserService.
func (u *UserServiceImpl) FindByUsername(username string) (userdata response.UserResponse, err error) {
	var userReponse response.UserResponse

	result, err := u.UserRepository.FindByUsername(username)
	// helper.ErrorPanic(err)

	if err != nil {
		userServicePrintln(err)
		return userReponse, err
	}

	userReponse = response.UserResponse{
		ID:       int(result.ID),
		Name:     result.Name,
		Username: result.Username,
		Age:      int(result.Age),
		RoleID:   int(*result.RoleID),
		Email:    result.Email,
		Phone:    result.Phone,
		Posts:    result.Posts,
	}

	return userReponse, nil
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(UserId int) {
	u.UserRepository.Delete(UserId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() (userList []response.UserResponse, err error) {

	var users = []response.UserResponse{}
	result, err := u.UserRepository.FindAll()
	// helper.ErrorPanic(err)

	if err != nil {
		userServicePrintln(err)
		return users, err
	}

	for _, v := range result {
		found := response.UserResponse{
			ID:       int(v.ID),
			Name:     v.Name,
			Username: v.Username,
			Age:      int(v.Age),
			RoleID:   int(*v.RoleID),
			Email:    v.Email,
			Phone:    v.Phone,
			Posts:    v.Posts,
		}

		users = append(users, found)
	}

	return users, nil
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(userId int) (userdata response.UserResponse, err error) {
	var userReponse response.UserResponse

	result, findErr := u.UserRepository.FindById(userId)
	// helper.ErrorPanic(findErr)

	if findErr != nil {
		userServicePrintln(findErr)
		return userReponse, findErr
	}

	userReponse = response.UserResponse{
		ID:       int(result.ID),
		Name:     result.Name,
		Username: result.Username,
		Age:      int(result.Age),
		RoleID:   int(*result.RoleID),
		Email:    result.Email,
		Phone:    result.Phone,
		Posts:    result.Posts,
	}

	return userReponse, nil
}

// Update implements UserService.
func (u *UserServiceImpl) Update(id int, user request.UpdateUserRequest) (userdata response.UserResponse, err error) {
	var userReponse response.UserResponse

	validateErr := u.Validate.Struct(user)

	if validateErr != nil {
		userServicePrintln(validateErr)
		return userReponse, validateErr
	}

	found, err := u.UserRepository.FindById(id)
	// helper.ErrorPanic(err)

	if err != nil {
		userServicePrintln(err)
		return userReponse, err
	}

	found.Name = user.Name
	// found.Username = user.Username
	found.Age = user.Age
	found.Email = user.Email
	found.Phone = user.Phone

	result, resultErr := u.UserRepository.Update(found)
	// helper.ErrorPanic(resultErr)

	if resultErr != nil {
		userServicePrintln(resultErr)
		return userReponse, resultErr
	}

	userReponse = response.UserResponse{
		ID:       int(result.ID),
		Name:     result.Name,
		Username: result.Username,
		Age:      int(result.Age),
		RoleID:   int(*result.RoleID),
		Email:    result.Email,
		Phone:    result.Phone,
		Posts:    result.Posts,
	}

	return userReponse, nil

}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
