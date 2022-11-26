package handler

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ibnumei/go-ms-playground/internal/app/domain"
)

type UserService interface {
	Register(context.Context, domain.User) (string, error)
}

type UserHandler struct {
	userService UserService 
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh UserHandler) Register (ctx *gin.Context) {
	var userBody domain.User
	fmt.Println(userBody)
	token, err := uh.userService.Register(ctx.Request.Context(), userBody)
	if err != nil{
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}