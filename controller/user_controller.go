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

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	id, err1 := strconv.ParseInt(userIdParam, 10, 64)
	if err1 != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return id, nil
}
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

func Get(ctx *gin.Context) {
	id, err := getUserId(ctx.Param("user_id"))
	if err != nil {
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

func Update(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
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

func Delete(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
