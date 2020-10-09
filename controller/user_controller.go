package controller

import (
	"fmt"
	"github.com/ghifar/bookstore-users-api/domain/users"
	"github.com/ghifar/bookstore-users-api/services"
	"github.com/ghifar/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user users.User

	//validate json format
	if err := user.ValidateJson(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	//Calling create user service
	res, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	fmt.Println("Object", user)
	ctx.JSON(http.StatusCreated, res)
}

func GetUser(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err1 != nil {
		err := errors.NewBadRequestError("invalid user id")
		ctx.JSON(err.Status, err)
		return
	}

	user, err2 := services.GetUser(id)
	if err2 != nil {
		ctx.JSON(err2.Status, err2)
		return
	}
	ctx.JSON(http.StatusOK, user)
	return

}

func FindUser(ctx *gin.Context) {

}

func UpdateUser(ctx *gin.Context) {
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Id should be a number")
		ctx.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := ctx.Request.Method == http.MethodPatch

	res, err := services.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
	return
}
