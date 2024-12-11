package main

import (
	"crud-postgres-orm/config"
	controller "crud-postgres-orm/controller/user"

	"crud-postgres-orm/model"
	repository "crud-postgres-orm/repository/user"
	"crud-postgres-orm/router"
	service "crud-postgres-orm/service/user"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func main() {
	db := config.ConnectDb()
	db.Table("users").AutoMigrate(&model.User{})
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, validate)
	userController := controller.NewUserController(userService, validate, trans)

	routes := router.NewRouter(userController)

	routes.Run(":8080")
}
