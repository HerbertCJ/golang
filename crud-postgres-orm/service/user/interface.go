package service

import (
	"crud-postgres-orm/data/request"
	"crud-postgres-orm/data/response"
)

type UserService interface {
	GetAll() []response.UserResponse
	GetById(id uint) (response.UserResponse, error)
	Delete(id uint)
	Create(user request.UserCreateRequest) error
	Update(user request.UserUpdateRequest, id uint) error
}
