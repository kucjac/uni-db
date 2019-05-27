package sqliteconv

import (
	"github.com/mattn/go-sqlite3"
	"github.com/neuronlabs/uni-db"
)

var defaultSQLiteErrorMap map[interface{}]unidb.Error = map[interface{}]unidb.Error{
	sqlite3.ErrWarning: unidb.ErrWarning,

	sqlite3.ErrNotFound: unidb.ErrNoResult,

	sqlite3.ErrCantOpen: unidb.ErrConnection,
	sqlite3.ErrNotADB:   unidb.ErrConnection,

	sqlite3.ErrMismatch: unidb.ErrDataException,

	sqlite3.ErrConstraint:           unidb.ErrIntegrityConstraintViolation,
	sqlite3.ErrConstraintCheck:      unidb.ErrCheckViolation,
	sqlite3.ErrConstraintForeignKey: unidb.ErrForeignKeyViolation,
	sqlite3.ErrConstraintUnique:     unidb.ErrUniqueViolation,
	sqlite3.ErrConstraintNotNull:    unidb.ErrNotNullViolation,
	sqlite3.ErrConstraintPrimaryKey: unidb.ErrUniqueViolation,

	sqlite3.ErrProtocol: unidb.ErrTxState,

	sqlite3.ErrRange: unidb.ErrInvalidSyntax,
	sqlite3.ErrError: unidb.ErrInvalidSyntax,

	sqlite3.ErrAuth: unidb.ErrAuthorizationFailed,

	sqlite3.ErrPerm: unidb.ErrInsufficientPrivilege,

	sqlite3.ErrFull: unidb.ErrInsufficientResources,

	sqlite3.ErrTooBig: unidb.ErrProgramLimitExceeded,

	sqlite3.ErrNoLFS: unidb.ErrSystemError,

	sqlite3.ErrInternal: unidb.ErrInternalError,
}
