package repository

import "go-crud/src/model"

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (user model.User, err error)
	FindAll() (user []model.User, err error)
}
