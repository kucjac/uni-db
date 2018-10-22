package sqliteconv

import (
	"github.com/kucjac/uni-db"
	"github.com/mattn/go-sqlite3"
)

var defaultSQLiteErrorMap map[interface{}]unidb.Error = map[interface{}]unidb.Error{
	sqlite3.ErrWarning: unidb.ErrWarning,

	sqlite3.ErrNotFound: unidb.ErrNoResult,

	sqlite3.ErrCantOpen: unidb.ErrConnExc,
	sqlite3.ErrNotADB:   unidb.ErrConnExc,

	sqlite3.ErrMismatch: unidb.ErrDataException,

	sqlite3.ErrConstraint:           unidb.ErrIntegrConstViolation,
	sqlite3.ErrConstraintCheck:      unidb.ErrCheckViolation,
	sqlite3.ErrConstraintForeignKey: unidb.ErrForeignKeyViolation,
	sqlite3.ErrConstraintUnique:     unidb.ErrUniqueViolation,
	sqlite3.ErrConstraintNotNull:    unidb.ErrNotNullViolation,
	sqlite3.ErrConstraintPrimaryKey: unidb.ErrUniqueViolation,

	sqlite3.ErrProtocol: unidb.ErrInvalidTransState,

	sqlite3.ErrRange: unidb.ErrInvalidSyntax,
	sqlite3.ErrError: unidb.ErrInvalidSyntax,

	sqlite3.ErrAuth: unidb.ErrInvalidAuthorization,

	sqlite3.ErrPerm: unidb.ErrInsufficientPrivilege,

	sqlite3.ErrFull: unidb.ErrInsufficientResources,

	sqlite3.ErrTooBig: unidb.ErrProgramLimitExceeded,

	sqlite3.ErrNoLFS: unidb.ErrSystemError,

	sqlite3.ErrInternal: unidb.ErrInternalError,
}
