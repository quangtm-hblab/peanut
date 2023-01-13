package controller

import (
	"net/http"
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Usecase: usecase.NewUserUsecase(db),
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

func (c *UserController) CreateUser(ctx *gin.Context) {
	user := domain.User{}
	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(user)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, nil)
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginForm domain.LoginForm
	if !bindJSON(ctx, &loginForm) {
		return
	}
	tokenString, errRes := c.Usecase.Login(ctx, loginForm)
	if errRes != nil {
		if errRes.Code == "400" {
			ctx.JSON(http.StatusBadRequest,
				domain.Response{false, nil, errRes.DebugMessage},
			)
			return
		} else if errRes.Code == "500" {
			ctx.JSON(http.StatusInternalServerError,
				domain.Response{false, nil, errRes.DebugMessage},
			)
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString, "message": "Login success"})
}
