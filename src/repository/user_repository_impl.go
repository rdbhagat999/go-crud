package repository

import (
	"fmt"
	"go-crud/src/data/request"
	"go-crud/src/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func userRepositoryPrintln(err error) {
	fmt.Println("userRepositoryPrintln: " + err.Error())
}

// Login implements UserRepository.
func (u *UserRepositoryImpl) Login(user request.LoginUserRequest) (authUser model.User, err error) {
	var foundUser model.User

	err = u.Db.Model(&model.User{}).Preload(clause.Associations).Find(&foundUser, "username = ?", user.Username).Error
	// helper.ErrorPanic(err)

	if err != nil {
		userRepositoryPrintln(err)
	}

	return foundUser, err

}

// FindByUsername implements UserRepository.
func (u *UserRepositoryImpl) FindByUsername(username string) (authUser model.User, err error) {
	var foundUser model.User

	err = u.Db.Model(&model.User{}).Preload(clause.Associations).Find(&foundUser, "username = ?", username).Error
	// helper.ErrorPanic(err)

	if err != nil {
		userRepositoryPrintln(err)
	}

	return foundUser, err

}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User

	err := u.Db.Select("Tags", "Posts").Where("id=?", userId).Delete(&user).Error
	// helper.ErrorPanic(result.Error)
	if err != nil {
		userRepositoryPrintln(err)
	}

}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() (users []model.User, err error) {
	var foundUsers []model.User
	// result := u.Db.Find(&foundUsers)
	// err = u.Db.Model(&model.User{}).Preload("Posts").Preload("Tags").Find(&foundUsers).Error
	err = u.Db.Model(&model.User{}).Preload(clause.Associations).Find(&foundUsers).Error

	if err != nil {
		userRepositoryPrintln(err)
	}

	return foundUsers, err
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId int) (user model.User, err error) {
	var foundUser model.User
	// result := u.Db.First(&foundUser, userId)
	// err = u.Db.Model(&model.User{}).Preload("Posts").Preload("Tags").Find(&foundUser, userId).Error
	err = u.Db.Model(&model.User{}).Preload(clause.Associations).Find(&foundUser, userId).Error

	if err != nil {
		userRepositoryPrintln(err)
		return foundUser, err
	}

	// if foundUser.ID == 0 {
	// 	return foundUser, errors.New("user not found")
	// }

	return foundUser, err
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) (us model.User, err error) {
	err = u.Db.Create(&user).Error
	// helper.ErrorPanic(err)

	if err != nil {
		userRepositoryPrintln(err)
		return model.User{}, err
	}

	return u.FindById(int(user.ID))
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) (us model.User, err error) {
	var updateUser = request.UpdateUserRequest{
		// Id:   user.Id,
		Name: user.Name,
		// Username: user.Username,
		Email: user.Email,
		Age:   int(user.Age),
		Phone: user.Phone,
	}

	err = u.Db.Model(&user).Updates(updateUser).Error
	// helper.ErrorPanic(err)

	if err != nil {
		userRepositoryPrintln(err)
		return model.User{}, err
	}

	return u.FindById(int(user.ID))

}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
