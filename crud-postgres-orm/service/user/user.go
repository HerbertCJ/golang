package service

import (
	"crud-postgres-orm/data/request"
	"crud-postgres-orm/data/response"
	"crud-postgres-orm/model"
	repository "crud-postgres-orm/repository/user"
	"errors"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

func (u *UserServiceImpl) GetById(id uint) response.UserResponse {
	user, err := u.UserRepository.GetById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.UserResponse{}
	}

	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	return userResponse
}

func (u *UserServiceImpl) Update(user request.UserUpdateRequest) error {
	_, err := u.UserRepository.GetById(user.Id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}

	userModel := model.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	u.UserRepository.Update(userModel)
	return nil
}
