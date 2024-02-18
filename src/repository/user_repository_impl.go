package repository

import (
	"go-crud/src/data/request"
	"go-crud/src/helper"
	"go-crud/src/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User

	result := u.Db.Model(&user).Delete(userId)
	// result := t.Db.Where("id=?", tagId).Delete(&tag)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() (users []model.User, err error) {
	var foundUsers []model.User
	// result := u.Db.Find(&foundUsers)
	err = u.Db.Model(&model.User{}).Preload("Tags").Find(&foundUsers).Error
	return foundUsers, err
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId int) (user model.User, err error) {
	var foundUser model.User
	// result := u.Db.First(&foundUser, userId)
	err = u.Db.Model(&model.User{}).Preload("Tags").Find(&foundUser, userId).Error
	return foundUser, err
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		// Id:   user.Id,
		Name: user.Name,
		// UserName: user.UserName,
		Email: user.Email,
		Age:   user.Age,
		Phone: user.Phone,
	}

	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
