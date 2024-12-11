package controller

import (
	"crud-postgres-orm/data/request"
	"crud-postgres-orm/data/response"
	"crud-postgres-orm/helper"
	service "crud-postgres-orm/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
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
	helper.ErrorPanic(err)

	userResponse := u.userService.GetById(uint(id))

	var resp response.WebResponse

	if userResponse.Id == 0 {
		resp = response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Ok",
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
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	u.userService.Create(createUserRequest)
	resp := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UserUpdateRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	updateUserRequest.Id = uint(id)

	u.userService.Update(updateUserRequest)
	resp := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}
