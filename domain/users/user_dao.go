package users

import (
	"fmt"
	"github.com/ghifar/bookstore-users-api/datasources/mysql/users_db"
	"github.com/ghifar/bookstore-users-api/logger"
	"github.com/ghifar/bookstore-users-api/utils/dateUtils"
	"github.com/ghifar/bookstore-users-api/utils/errors"
	"github.com/ghifar/bookstore-users-api/utils/mysqlUtils"
)

const (
	Q_INSERT_USER         = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	Q_GET_USER            = "SELECT id, email, first_name, last_name, date_created, status FROM users WHERE id=?;"
	Q_UPDATE_USER         = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	Q_DELETE_USER         = "DELETE FROM users WHERE id=?;"
	Q_FIND_USER_BY_STATUS = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(Q_GET_USER)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//note: returns single row only
	getRes := stmt.QueryRow(user.Id)
	if err := getRes.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.DateCreated, &user.Status); err != nil {
		return mysqlUtils.SqlErrorParser(err)
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

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, dateUtils.GetNowDbFormat(), user.Status, user.Password)
	if err != nil {
		return mysqlUtils.SqlErrorParser(err)
	}

	userId, err := insertRes.LastInsertId()
	if err != nil {
		return mysqlUtils.SqlErrorParser(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(Q_UPDATE_USER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysqlUtils.SqlErrorParser(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(Q_DELETE_USER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	deleteRes, err := stmt.Exec(user.Id)
	if err != nil {
		return mysqlUtils.SqlErrorParser(err)
	}
	_, err = deleteRes.RowsAffected()
	if err != nil {
		return mysqlUtils.SqlErrorParser(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(Q_FIND_USER_BY_STATUS)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	res := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status);
			err != nil {
			return nil, mysqlUtils.SqlErrorParser(err)
		}
		res = append(res, user)
	}

	if len(res) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching with status %s", status))
	}
	return res, nil
}
