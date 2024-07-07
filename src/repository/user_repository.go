package repository

import (
	"go-crud/src/data/request"
	"go-crud/src/model"
)

type UserRepository interface {
	Save(user model.User) (savedUser model.User, err error)
	Login(user request.LoginUserRequest) (authUser model.User, err error)
	Update(user model.User) (updatedUser model.User, err error)
	Delete(userId int)
	FindById(userId int) (foundUser model.User, err error)
	FindByUsername(username string) (foundUser model.User, err error)
	FindAll() (users []model.User, err error)
}
