package mysqlUtils

import (
	"fmt"
	"github.com/ghifar/bookstore-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	NO_ROWS = "no rows in result set"
)

func SqlErrorParser(errs error) *errors.RestErr {
	errTemp, isSQLError := errs.(*mysql.MySQLError)

	if !isSQLError {
		fmt.Println(errs)
		if strings.Contains(errs.Error(), NO_ROWS) {
			return errors.NewNotFoundError(fmt.Sprintf("Couldn't find given id. %s", errs.Error()))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Server error: %s", errs.Error()))
	}

	fmt.Println(errTemp)
	switch errTemp.Number{
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("Duplicate key: %s", errTemp.Message))
	default:
		return errors.NewBadRequestError(fmt.Sprintf("SQL Error: %s", errTemp.Message))
	}
}
