package services

import (
	"github.com/ghifar/bookstore-users-api/domain/users"
	"github.com/ghifar/bookstore-users-api/utils"
	"github.com/ghifar/bookstore-users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (UserService *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	res := &users.User{Id: userId}
	if err := res.Get(); err != nil {
		return nil, err
	}
	return res, nil
}

func (UserService *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateField(); err != nil {
		//panic(err)
		return nil, err
	}

	user.Password = utils.GetMd5(user.Password)
	user.Status = "active"
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (UserService *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	//find the user
	curr, err := UserService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			curr.FirstName = user.FirstName
		}
		if user.LastName != "" {
			curr.LastName = user.LastName
		}
		if user.Email != "" {
			curr.Email = user.Email
		}

	} else {
		curr.FirstName = user.FirstName
		curr.LastName = user.LastName
		curr.Email = user.Email
	}

	if err := curr.Update(); err != nil {
		return nil, err
	}
	return curr, nil
}

func (UserService *userService) DeleteUser(userId int64) *errors.RestErr {
	res := &users.User{Id: userId}
	return res.Delete()
}

func (UserService *userService) Search(status string) (users.Users, *errors.RestErr) {
	userDao := &users.User{}
	return userDao.FindByStatus(status)
}
