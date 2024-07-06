package repository

import "go-crud/src/model"

type UserRepository interface {
	Save(user model.User) (u model.User, err error)
	Update(user model.User) (u model.User, err error)
	Delete(userId int)
	FindById(userId int) (user model.User, err error)
	FindAll() (users []model.User, err error)
}
