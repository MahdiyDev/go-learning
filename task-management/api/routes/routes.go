package routes

import (
	"tm/lib"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	userRoutes UserRoutes
}

func (u *Routes) Setup(router *gin.Engine, conn *lib.DatabaseConnection) {
	u.userRoutes = UserRoutes{}

	u.userRoutes.Setup(router, conn)
}
