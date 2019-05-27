package sqliteconv

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"github.com/neuronlabs/uni-db"
)

// SQLiteConverter is ErrorConverter interface implementation
// for sqlite3 database.
//
type SQLiteConverter struct {
	errorMap map[interface{}]unidb.Error
}

// Convert converts the provided error into *Error type.
// It is method that implements ErrorConverter Interface
func (r *SQLiteConverter) Convert(err error) *unidb.Error {
	// Check if the error is of '*sqlite3.Error' type
	sqliteErr, ok := err.(sqlite3.Error)
	if !ok {
		// if not check sql errors
		if err == sql.ErrNoRows {
			return unidb.ErrNoResult.NewWithError(err)
		} else if err == sql.ErrTxDone {
			return unidb.ErrTxDone.NewWithError(err)
		}
		return unidb.ErrUnspecifiedError.NewWithError(err)
	}

	var dbError unidb.Error
	// Check if Error.ExtendedCode is in recogniser
	dbError, ok = r.errorMap[sqliteErr.ExtendedCode]
	if ok {
		return dbError.NewWithError(err)
	}

	// otherwise check if Error.Code is in the recogniser
	dbError, ok = r.errorMap[sqliteErr.Code]
	if ok {
		return dbError.NewWithError(err)
	}

	// if no error is specified return Unspecified Error
	return unidb.ErrUnspecifiedError.NewWithError(err)
}

func New() *SQLiteConverter {
	return &SQLiteConverter{errorMap: defaultSQLiteErrorMap}
}
