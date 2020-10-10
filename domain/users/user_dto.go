package users

import (
	"github.com/ghifar/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) ValidateField() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}

func (user *User) ValidateJson(ctx *gin.Context) *errors.RestErr {
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Println("Error validation")
		err := errors.NewBadRequestError("Invalid json body")
		return err
	}
	return nil
}
