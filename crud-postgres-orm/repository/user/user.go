package repository

import (
	"crud-postgres-orm/helper"
	"errors"

	"crud-postgres-orm/model"

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

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}

		return model.User{}, result.Error
	} else {
		return user, nil
	}
}

func (u *UserRepositoryImpl) Create(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) Update(user model.User) {
	result := u.Db.Model(&user).Updates(user)
	helper.ErrorPanic(result.Error)
}
