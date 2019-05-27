package unidb

var (
	// ErrWarning is a warning message
	ErrWarning = Error{ID: 1, Title: "Warning"}

	// ErrNoResult used as a replacement for ErrNoRows - for non-sql databases
	ErrNoResult = Error{ID: 2, Title: "No Result"}

	// ErrConnection is an error thrown when there are problems with the database connection
	ErrConnection = Error{ID: 3, Title: "Connection exception"}

	// ErrCardinalityViolation is an error thrown when the cardinality of results fails
	// i.e. the result of subquery contains 4 fields whereas in the root query requires only 2
	ErrCardinalityViolation = Error{ID: 4, Title: "Cardinality violation"}

	// ErrDataException data exception is the data format and integration error
	ErrDataException = Error{ID: 5, Title: "Data Exception"}

	// Violations

	// ErrIntegrityConstraintViolation is an error thrown when integrity constarint violation occurs
	ErrIntegrityConstraintViolation = Error{ID: 6, Title: "Integrity constraint violation"}

	// ErrRestrictViolation is an error thrown when the restricted violation is thrown
	ErrRestrictViolation = Error{ID: 7, Title: "Restrict violation"}

	// ErrNotNullViolation thrown when the NOT NULL restriction is violated
	ErrNotNullViolation = Error{ID: 8, Title: "Not null violation"}

	// ErrForeignKeyViolation error thrown when the Foreign Key is violated
	ErrForeignKeyViolation = Error{ID: 9, Title: "Foreign-Key violation"}

	// ErrUniqueViolation thrown when the restriction on the Uniqness is violated
	ErrUniqueViolation = Error{ID: 10, Title: "Unique violation"}

	// ErrCheckViolation thrown when any check on the data value is violated
	ErrCheckViolation = Error{ID: 11, Title: "Check violation"}

	// Transactions

	// ErrTxState is an error thrown when the transaction is on invalid state
	ErrTxState = Error{ID: 12, Title: "Invalid transaction state"}

	// ErrTxTermination thrown when the transaction is already terminated
	ErrTxTermination = Error{ID: 13, Title: "Invalid transaction termination"}
	// ErrTxRollback thrown when the current transaction is already Rollbacked
	ErrTxRollback = Error{ID: 14, Title: "Transaction Rollback"}

	// ErrTxNotFound thrown when the transaction is not found
	ErrTxNotFound = Error{ID: 15, Title: "Transaction not found"}

	// ErrTxBeginInProgress thrown when the transaction Begin phase is still in progress
	ErrTxBeginInProgress = Error{ID: 16, Title: "Transaction in progress"}

	// ErrTxDone thrown when the transaction is already Done while trying to Commit the changes
	ErrTxDone = Error{ID: 15, Title: "Transaction done"}

	// ErrAuthorizationFailed thrown when the authorization process failed
	ErrAuthorizationFailed = Error{ID: 16, Title: "Invalid Authorization Specification"}

	// ErrAuthenticationFailed thrown when authentication failed while connecting to the database
	ErrAuthenticationFailed = Error{ID: 17, Title: "Authentication failed"}

	// ErrInvalidSchemaName thrown when the provided schema name doesn't exists
	ErrInvalidSchemaName = Error{ID: 18, Title: "Invalid Schema Name"}

	// ErrInvalidCatalogName thrown when the catalog name is invalid
	ErrInvalidCatalogName = Error{ID: 19, Title: "Invalid Catalog Name"}

	// ErrInvalidResourceName thrown when the resource name is not valid
	ErrInvalidResourceName = Error{ID: 28, Title: "Invalid Resource name"}

	// ErrInvalidSyntax thrown for an invalid syntax provided
	ErrInvalidSyntax = Error{ID: 20, Title: "Syntax Error"}

	// ErrInsufficientPrivilege thrown when the role has not sufficient privileges
	ErrInsufficientPrivilege = Error{ID: 21, Title: "Insufficient Privilege"}

	// ErrInsufficientResources thrown when the resources of the database is not sufficient for given query
	ErrInsufficientResources = Error{ID: 22, Title: "Insufficient Resources"}

	// ErrProgramLimitExceeded thrown when the database configuration limits are exceeded
	ErrProgramLimitExceeded = Error{ID: 23, Title: "Program Limit Exceeded"}

	// ErrSystemError thrown when there are errors with the system (i.e. I/O errors)
	ErrSystemError = Error{ID: 24, Title: "System error"}

	// ErrInternalError thrown on internal errors
	ErrInternalError = Error{ID: 25, Title: "Internal error"}

	// ErrUnspecifiedError - all other errors not included in this division
	ErrUnspecifiedError = Error{ID: 26, Title: "Unspecified error"}

	// ErrTimedOut when the query (not connection) is timed out
	ErrTimedOut = Error{ID: 27, Title: "Timed out"}

	// ErrReplica is the error that is thrown when there is a problem with replications
	ErrReplica = Error{ID: 29, Title: "Replica('s) error"}

	// ErrShutdown thrown when the database server is shutting down
	ErrShutdown = Error{ID: 32, Title: "Shutting down"}
)

var prototypeMap = map[uint]Error{
	uint(1):  ErrWarning,
	uint(2):  ErrNoResult,
	uint(3):  ErrConnection,
	uint(4):  ErrCardinalityViolation,
	uint(5):  ErrDataException,
	uint(6):  ErrIntegrityConstraintViolation,
	uint(7):  ErrRestrictViolation,
	uint(8):  ErrNotNullViolation,
	uint(9):  ErrForeignKeyViolation,
	uint(10): ErrUniqueViolation,
	uint(11): ErrCheckViolation,
	uint(12): ErrTxState,
	uint(13): ErrTxTermination,
	uint(14): ErrTxRollback,
	uint(15): ErrTxDone,
	uint(16): ErrAuthorizationFailed,
	uint(17): ErrAuthenticationFailed,
	uint(18): ErrInvalidSchemaName,
	uint(19): ErrInvalidCatalogName,
	uint(20): ErrInvalidSyntax,
	uint(21): ErrInsufficientPrivilege,
	uint(22): ErrInsufficientResources,
	uint(23): ErrProgramLimitExceeded,
	uint(24): ErrSystemError,
	uint(25): ErrInternalError,
	uint(26): ErrUnspecifiedError,
	uint(27): ErrTimedOut,
	uint(28): ErrInvalidResourceName,
	uint(29): ErrReplica,
	uint(30): ErrTxNotFound,
	uint(32): ErrShutdown,
}
