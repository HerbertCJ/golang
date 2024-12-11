package main

import (
	"crud-postgres-orm/config"
	controller "crud-postgres-orm/controller/user"

	"crud-postgres-orm/model"
	repository "crud-postgres-orm/repository/user"
	"crud-postgres-orm/router"
	service "crud-postgres-orm/service/user"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.ConnectDb()
	db.Table("users").AutoMigrate(&model.User{})
	validate := validator.New()

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, validate)
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)

	routes.Run(":8080")
}
