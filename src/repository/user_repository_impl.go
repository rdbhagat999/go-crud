package repository

import (
	"go-crud/src/data/request"
	"go-crud/src/helper"
	"go-crud/src/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User

	result := u.Db.Select("Tags", "Posts").Model(&user).Delete(userId)
	// result := u.Db.Select("Tags", "Posts").Where("id=?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() (users []model.User, err error) {
	var foundUsers []model.User
	// result := u.Db.Find(&foundUsers)
	// err = u.Db.Model(&model.User{}).Preload("Posts").Preload("Tags").Find(&foundUsers).Error
	err = u.Db.Model(&model.User{}).Preload(clause.Associations).Find(&foundUsers).Error
	return foundUsers, err
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId int) (user model.User, err error) {
	var foundUser model.User
	// result := u.Db.First(&foundUser, userId)
	// err = u.Db.Model(&model.User{}).Preload("Posts").Preload("Tags").Find(&foundUser, userId).Error
	err = u.Db.Model(&model.User{}).Preload(clause.Associations).Find(&foundUser, userId).Error
	return foundUser, err
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) (us model.User, err error) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)

	return u.FindById(user.ID)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) (us model.User, err error) {
	var updateUser = request.UpdateUserRequest{
		// Id:   user.Id,
		Name: user.Name,
		// Username: user.Username,
		Email: user.Email,
		Age:   user.Age,
		Phone: user.Phone,
	}

	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)

	return u.FindById(user.ID)

}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
