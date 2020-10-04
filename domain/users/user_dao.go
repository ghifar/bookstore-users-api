package users

import (
	"fmt"
	"github.com/ghifar/bookstore-users-api/datasources/mysql/users_db"
	"github.com/ghifar/bookstore-users-api/utils/dateUtils"
	"github.com/ghifar/bookstore-users-api/utils/errors"
	"strings"
)

const (
	Q_INSERT_USER = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	Q_GET_USER    = "SELECT id, email, first_name, last_name, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(Q_GET_USER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	getRes := stmt.QueryRow(user.Id)
	if err := getRes.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errors.NO_ROWS) {
			return errors.NewBadRequestError(fmt.Sprintf("Couldn't find id: %d", user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to get user %d : %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	//note: prepare has a better performance and advantages than directly execute query from client like so client.Exec(query, ...param).
	stmt, err := users_db.Client.Prepare(Q_INSERT_USER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, dateUtils.GetNowString())
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError(fmt.Sprintf("Email already registered: %s", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	userId, err := insertRes.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
