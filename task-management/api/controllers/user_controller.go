package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tm/entities"
	"tm/lib"
	"tm/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `json:"username" validate:"required,min=3,max=20"`
	Age  int    `json:"age" validate:"required,gte=1"`
}

type UserController struct {
	userService services.UserService
}

func (u *UserController) New(conn *lib.DatabaseConnection) {
	u.userService.New(conn)
}

func (u UserController) GetAllUsers(ctx *gin.Context) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "some thing went wrong"},
		)
		return
	}
	if len(users) == 0 {
		ctx.JSON(
			http.StatusNotFound,
			map[string]string{"message": "users not found"},
		)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u UserController) GetOneUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("UserController GetOneUsers error: %v\n", err)
		return
	}
	user, err := u.userService.GetOneUsers(parsedID)
	validate := validator.New()
	valErr := validate.Struct(user)
	if valErr != nil {
		fmt.Println(valErr)
		ctx.JSON(
			http.StatusBadRequest,
			map[string]string{"message": "user must be not empty"},
		)
		return
	}
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "some thing went wrong"},
		)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u UserController) CreateUser(ctx *gin.Context) {
	userEntity := entities.UserEntity{}
	if err := ctx.BindJSON(&userEntity); err != nil {
		fmt.Printf("UserController CreateUser error: %v\n", err)
		return
	}
	validate := validator.New()
	err := validate.Struct(userEntity)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			map[string]string{"message": "user must be not empty"},
		)
		return
	}
	users, err := u.userService.CreateUser(userEntity)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "some thing went wrong"},
		)
		return
	}
	ctx.JSON(http.StatusOK, users)
}
