package users

import (
	"github.com/ghifar/bookstore-users-api/datasources/mysql/users_db"
	"github.com/ghifar/bookstore-users-api/utils/dateUtils"
	"github.com/ghifar/bookstore-users-api/utils/errors"
	"github.com/ghifar/bookstore-users-api/utils/mysqlUtils"
)

const (
	Q_INSERT_USER = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	Q_GET_USER    = "SELECT id, email, first_name, last_name, date_created FROM users WHERE id=?;"
	Q_UPDATE_USER = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	Q_DELETE_USER = "DELETE FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(Q_GET_USER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//note: returns single row only
	getRes := stmt.QueryRow(user.Id)
	if err := getRes.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.DateCreated); err != nil {
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

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, dateUtils.GetNowString())
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
