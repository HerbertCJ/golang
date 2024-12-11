package controller

import (
	"crud-postgres-orm/data/request"
	"crud-postgres-orm/data/response"
	"crud-postgres-orm/helper"
	service "crud-postgres-orm/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService service.UserService
	Validate    *validator.Validate
	trans       ut.Translator
}

func NewUserController(userService service.UserService, validate *validator.Validate, trans ut.Translator) *UserController {
	return &UserController{userService: userService, Validate: validate, trans: trans}
}

func (u *UserController) GetAll(ctx *gin.Context) {
	usersResponse := u.userService.GetAll()

	resp := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   usersResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserController) GetById(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)

	if err != nil {
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Data:    nil,
			Message: "Invalid user ID",
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	userResponse, err := u.userService.GetById(uint(id))

	var resp response.WebResponse

	if err != nil {
		resp = response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Error",
			Data:    nil,
			Message: "User not found",
		}
	} else {
		resp = response.WebResponse{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   userResponse,
		}
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	u.userService.Delete(uint(id))

	resp := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.UserCreateRequest{}

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Data:    nil,
			Message: "Campos inválidos",
			Errors:  helper.TranslateError(err, u.trans),
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := u.Validate.Struct(createUserRequest); err != nil {
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Data:    nil,
			Message: "Bad request",
			Errors:  helper.TranslateError(err, u.trans),
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err := u.userService.Create(createUserRequest)

	if err != nil {
		resp := response.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Error",
			Data:    nil,
			Message: "Some error occurred, try again later!",
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "User created",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UserUpdateRequest{}
	if err := ctx.ShouldBindJSON(&updateUserRequest); err != nil {
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Data:    nil,
			Message: "Campos inválidos",
			Errors:  helper.TranslateError(err, u.trans),
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := u.Validate.Struct(updateUserRequest); err != nil {
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Data:    nil,
			Message: "Bad request",
			Errors:  helper.TranslateError(err, u.trans),
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)

	if err != nil {
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Data:    nil,
			Message: "Invalid user ID",
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = u.userService.Update(updateUserRequest, uint(id))

	if err != nil {
		resp := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Error",
			Data:    nil,
			Message: err.Error(),
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	resp := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}
