package service

import (
	"crud-postgres-orm/data/request"
	"crud-postgres-orm/data/response"
	"crud-postgres-orm/model"
	repository "crud-postgres-orm/repository/user"
	"errors"

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

func (u *UserServiceImpl) Create(user request.UserCreateRequest) error {
	err := u.Validate.Struct(user)

	if err != nil {
		return err
	}

	userModel := model.User{
		Username: user.Username,
		Email:    user.Email,
	}

	u.UserRepository.Create(userModel)
	return nil
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

func (u *UserServiceImpl) GetById(id uint) (response.UserResponse, error) {
	user, err := u.UserRepository.GetById(id)

	if err != nil {
		return response.UserResponse{}, err
	}

	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	return userResponse, nil
}

func (u *UserServiceImpl) Update(user request.UserUpdateRequest, id uint) error {
	_, err := u.UserRepository.GetById(id)

	if err != nil {
		return errors.New("user not found")
	}

	userModel := model.User{
		Id:       id,
		Username: user.Username,
		Email:    user.Email,
	}

	u.UserRepository.Update(userModel)
	return nil
}
