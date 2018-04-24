package mysqlconv

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/kucjac/uni-db"
)

// MySQLConverter is a Converter interface implementation
// The Converter can convert provided error into *Error with specific logic.
// Check the Convert method documentation for more information on how it distinguish given error
type MySQLConverter struct {
	// codeMap puts an error code or sqlstate into map and returns unidb.Error prototype
	codeMap map[interface{}]unidb.Error

	// sqlStateMap is a helper map that recognises the sqlstate for given error code
	sqlStateMap map[uint16]string
}

// Convert converts provided 'err' error into *unidb.Error type.
// With this method MySQLConverter implements ErrorConverter interface.
// Convert distinguish  and convert specific error of types sql.Err*, mysql.Err*,
// and *mysql.MySQLError. If an error is of different type it returns new entity of
// unidb.ErrUnspecifiedError
// If the error is of *mysql.MySQLError type the method checks its code.
// If the code matches with internal code map it returns proper entity of *unidb.Error.
// If the code does not exists in the code map, the method gets sqlstate for given code
// and checks if this sqlstate is in the code map.
// If the sqlstate does not exists in the code map, the first two numbers from the sqlstate
// are being checked in the codeMap as a 'sqlstate class'.
// If not found Convert returns new entity for unidb.UnspecifiedError
func (m *MySQLConverter) Convert(err error) *unidb.Error {
	// Check whether the given error is of *mysql.MySQLError
	mySQLErr, ok := err.(*mysql.MySQLError)
	if !ok {
		// Otherwise check if it sql.Err* or other errors from mysql package
		switch err {
		case mysql.ErrInvalidConn, mysql.ErrNoTLS, mysql.ErrOldProtocol,
			mysql.ErrMalformPkt:
			return unidb.ErrConnExc.NewWithError(err)
		case sql.ErrNoRows:
			return unidb.ErrNoResult.NewWithError(err)

		case sql.ErrTxDone:
			return unidb.ErrTxDone.NewWithError(err)

		default:
			return unidb.ErrUnspecifiedError.NewWithError(err)
		}
	}
	var dbErr unidb.Error

	// Check if Error Number is in recogniser
	dbErr, ok = m.codeMap[mySQLErr.Number]
	if ok {
		// Return if found
		return dbErr.NewWithError(err)
	}

	// Otherwise check if given sqlstate is in the codeMap
	sqlState, ok := m.sqlStateMap[mySQLErr.Number]
	if !ok || len(sqlState) != 5 {
		return unidb.ErrUnspecifiedError.NewWithError(err)
	}
	dbErr, ok = m.codeMap[sqlState]
	if ok {
		return dbErr.NewWithError(err)
	}

	// First two letter from sqlState represents error class
	// Check if class is in error map
	sqlStateClass := sqlState[0:2]
	dbErr, ok = m.codeMap[sqlStateClass]
	if ok {
		return dbErr.NewWithError(err)
	}

	return unidb.ErrUnspecifiedError.NewWithError(err)
}

// New creates new already inited MySQLConverter
func New() *MySQLConverter {
	return &MySQLConverter{
		codeMap:     mysqlErrMap,
		sqlStateMap: codeSQLState,
	}
}
