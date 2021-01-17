package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/cry999/pm-projects/pkg/domain/errors/common"
	"github.com/go-sql-driver/mysql"
)

const (
	formatErrDuplicateEntry = "Duplicate entry %s for key %s"
)

const (
	mysqlCodeDuplicateEntry = 1062
)

// HandleMySQLError ...
func HandleMySQLError(err error, entity, key string) error {
	if err == sql.ErrNoRows {
		return common.NotFoundError(entity, key)
	}
	switch err := err.(type) {
	case *mysql.MySQLError:
		fmt.Fprintf(os.Stderr, "code=%v,message=%v\n", err.Number, err.Message)
		switch err.Number {
		case mysqlCodeDuplicateEntry: // Duplicate key
			var key, val string
			if _, err := fmt.Sscanf(err.Message, formatErrDuplicateEntry, &val, &key); err != nil {
				return common.IllegalOperationError(err.Error())
			}
			return common.IllegalOperationError("%s: %s already exists", key, val)
		}
		return err
	}
	if err := errors.Unwrap(err); err != nil {
		return HandleMySQLError(err, entity, key)
	}
	return err
}
