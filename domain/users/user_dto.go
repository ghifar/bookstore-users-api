package users

import (
	"github.com/ghifar/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type User struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	DateCreated  string `json:"date_created"`
}

func (user *User) ValidateField() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}

func (user *User) ValidateJson(ctx *gin.Context)  *errors.RestErr{
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Println("Error validation")
		err := errors.NewBadRequestError("Invalid json body")
		return err
	}
	return nil
}
