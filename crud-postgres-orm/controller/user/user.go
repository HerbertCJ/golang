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

	resp := helper.WebResponseFormatter(http.StatusOK, "Ok", usersResponse, "Users found", nil)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(resp.Code, resp)
}

func (u *UserController) GetById(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)

	if err != nil {
		resp := helper.WebResponseFormatter(http.StatusBadRequest, "Error", nil, "Invalid user ID", nil)
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	userResponse, err := u.userService.GetById(uint(id))

	var resp response.WebResponse
	if err != nil {
		resp = helper.WebResponseFormatter(http.StatusNotFound, "Error", nil, "User not found", nil)
	} else {
		resp = helper.WebResponseFormatter(http.StatusOK, "Ok", userResponse, "User found", nil)
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(resp.Code, resp)
}

func (u *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	u.userService.Delete(uint(id))

	resp := helper.WebResponseFormatter(http.StatusOK, "Ok", nil, "User deleted", nil)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(resp.Code, resp)
}

func (u *UserController) Create(ctx *gin.Context) {
	var resp response.WebResponse
	createUserRequest := request.UserCreateRequest{}

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		resp = helper.WebResponseFormatter(http.StatusBadRequest, "Error", nil, "Campos inválidos", helper.TranslateError(err, u.trans))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	if err := u.Validate.Struct(createUserRequest); err != nil {
		resp = helper.WebResponseFormatter(http.StatusBadRequest, "Error", nil, "Bad request", helper.TranslateError(err, u.trans))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	err := u.userService.Create(createUserRequest)

	if err != nil {
		resp = helper.WebResponseFormatter(http.StatusInternalServerError, "Error", nil, "Some error occurred, try again later!", helper.TranslateError(err, u.trans))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	resp = helper.WebResponseFormatter(http.StatusOK, "Ok", nil, "User created", helper.TranslateError(err, u.trans))
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(resp.Code, resp)
}

func (u *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UserUpdateRequest{}
	var resp response.WebResponse

	if err := ctx.ShouldBindJSON(&updateUserRequest); err != nil {
		resp = helper.WebResponseFormatter(http.StatusBadRequest, "Error", nil, "Campos inválidos", helper.TranslateError(err, u.trans))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	if err := u.Validate.Struct(updateUserRequest); err != nil {
		resp = helper.WebResponseFormatter(http.StatusBadRequest, "Error", nil, "Bad request", helper.TranslateError(err, u.trans))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)

	if err != nil {
		resp = helper.WebResponseFormatter(http.StatusBadRequest, "Error", nil, "Invalid user ID", helper.TranslateError(err, u.trans))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	err = u.userService.Update(updateUserRequest, uint(id))

	if err != nil {
		resp = helper.WebResponseFormatter(http.StatusNotFound, "Error", nil, err.Error(), helper.TranslateError(err, u.trans))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(resp.Code, resp)
		return
	}

	resp = helper.WebResponseFormatter(http.StatusOK, "Ok", nil, "", helper.TranslateError(err, u.trans))
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(resp.Code, resp)
}
