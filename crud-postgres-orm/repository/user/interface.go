package repository

import (
	"crud-postgres-orm/model"
)

type UserRepository interface {
	GetAll() []model.User
	GetById(id uint) (model.User, error)
	Delete(id uint)
	Create(user model.User)
	Update(user model.User)
}
