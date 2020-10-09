package services

import (
	"github.com/ghifar/bookstore-users-api/domain/users"
	"github.com/ghifar/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateField(); err != nil {
		panic(err)
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	res := &users.User{Id: userId}
	if err := res.Get(); err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	//find the user
	curr, err := GetUser(user.Id)
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
