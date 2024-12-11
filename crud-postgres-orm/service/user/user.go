package service

import (
	"crud-postgres-orm/data/request"
	"crud-postgres-orm/data/response"
	"crud-postgres-orm/helper"
	"crud-postgres-orm/model"
	repository "crud-postgres-orm/repository/user"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) Create(user request.UserCreateRequest) {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.User{
		Username: user.Username,
		Email:    user.Email,
	}

	u.UserRepository.Create(userModel)
}

func (u *UserServiceImpl) Delete(id uint) {
	u.UserRepository.Delete(id)
}

func (u *UserServiceImpl) GetAll() []response.UserResponse {
	result := u.UserRepository.GetAll()

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:       value.Id,
			Username: value.Username,
			Email:    value.Email,
		}
		users = append(users, user)
	}

	return users
}

func (u *UserServiceImpl) GetById(id uint) response.UserResponse {
	user, err := u.UserRepository.GetById(id)
	helper.ErrorPanic(err)

	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	return userResponse
}

func (u *UserServiceImpl) Update(user request.UserUpdateRequest) {
	_, err := u.UserRepository.GetById(user.Id)
	helper.ErrorPanic(err)

	userModel := model.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	u.UserRepository.Update(userModel)
}
