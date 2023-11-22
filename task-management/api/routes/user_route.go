package routes

import (
	"tm/api/controllers"
	"tm/lib"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController controllers.UserController
}

func (u *UserRoutes) Setup(router *gin.Engine, conn *lib.DatabaseConnection) {
	u.userController.New(conn)

	router.GET("/user/", u.userController.GetAllUsers)
	router.GET("/user/:id", u.userController.GetOneUsers)
	router.POST("/user/", u.userController.CreateUser)
}
