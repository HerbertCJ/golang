package repository

import (
	"crud-postgres-orm/data/request"
	"crud-postgres-orm/helper"

	"crud-postgres-orm/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

func (u *UserRepositoryImpl) Delete(id uint) {
	result := u.Db.Delete(&model.User{}, id)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) GetAll() []model.User {
	var users []model.User
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)

	return users
}

func (u *UserRepositoryImpl) GetById(id uint) (model.User, error) {
	var user model.User
	result := u.Db.First(&user, id)

	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u *UserRepositoryImpl) Create(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UserUpdateRequest{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
