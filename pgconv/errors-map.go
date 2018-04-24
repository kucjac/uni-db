package pgconv

import (
	"github.com/kucjac/uni-db"
	"github.com/lib/pq"
)

var defaultPGErrorMap = map[interface{}]unidb.Error{

	// Class 01 - Warnings
	pq.ErrorClass("01"): unidb.ErrWarning,

	// Class 02 - No data
	pq.ErrorClass("02"):   unidb.ErrNoResult,
	pq.ErrorCode("P0002"): unidb.ErrNoResult,

	// Class 08 - Connection Exception
	pq.ErrorClass("08"): unidb.ErrConnExc,

	// Class 21 - Cardinality Violation
	pq.ErrorClass("21"): unidb.ErrCardinalityViolation,

	// Class 22 Data Exception
	pq.ErrorClass("22"): unidb.ErrDataException,

	// Class 23 Integrity Violation errors
	pq.ErrorClass("23"):   unidb.ErrIntegrConstViolation,
	pq.ErrorCode("23000"): unidb.ErrIntegrConstViolation,
	pq.ErrorCode("23001"): unidb.ErrRestrictViolation,
	pq.ErrorCode("23502"): unidb.ErrNotNullViolation,
	pq.ErrorCode("23503"): unidb.ErrForeignKeyViolation,
	pq.ErrorCode("23505"): unidb.ErrUniqueViolation,
	pq.ErrorCode("23514"): unidb.ErrCheckViolation,

	// Class 25 Invalid Transaction State
	pq.ErrorClass("25"): unidb.ErrInvalidTransState,

	// Class 28 Invalid Authorization Specification
	pq.ErrorCode("28000"): unidb.ErrInvalidAuthorization,
	pq.ErrorCode("28P01"): unidb.ErrInvalidPassword,

	// Class 2D Invalid Transaction Termination
	pq.ErrorCode("2D000"): unidb.ErrInvalidTransTerm,

	// Class 3F Invalid Schema Name
	pq.ErrorCode("3F000"): unidb.ErrInvalidSchemaName,

	// Class 40 - Transaciton Rollback
	pq.ErrorClass("40"): unidb.ErrTransRollback,

	// Class 42 - Invalid Syntax
	pq.ErrorClass("42"):   unidb.ErrInvalidSyntax,
	pq.ErrorCode("42501"): unidb.ErrInsufficientPrivilege,

	// Class 53 - Insufficient Resources
	pq.ErrorClass("53"): unidb.ErrInsufficientResources,

	// Class 54 - Program Limit Exceeded
	pq.ErrorClass("54"): unidb.ErrProgramLimitExceeded,

	// Class 58 - System Errors
	pq.ErrorClass("58"): unidb.ErrSystemError,

	// Class XX - Internal Error
	pq.ErrorClass("XX"): unidb.ErrInternalError,
}
