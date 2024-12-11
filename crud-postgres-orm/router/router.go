package router

import (
	controller "crud-postgres-orm/controller/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	userRouter := router.Group("/users")
	userRouter.GET("", userController.GetAll)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:id", userController.Update)
	userRouter.GET("/:id", userController.GetById)
	userRouter.DELETE("/:id", userController.Delete)

	return router
}
