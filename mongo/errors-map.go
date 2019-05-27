package mongo

import (
	"github.com/neuronlabs/uni-db"
	"go.mongodb.org/mongo-driver/mongo"
)

func defaultMapping() map[interface{}]unidb.Error {
	return map[interface{}]unidb.Error{
		// mongo predefined errors
		mongo.ErrClientDisconnected: unidb.ErrConnection,
		mongo.ErrEmptySlice:         unidb.ErrNoResult,
		mongo.ErrNilDocument:        unidb.ErrCheckViolation,
		mongo.ErrNoDocuments:        unidb.ErrNoResult,
		mongo.ErrWrongClient:        unidb.ErrConnection,

		// codes
		int32(1):  unidb.ErrInternalError,
		int32(2):  unidb.ErrCheckViolation,
		int32(3):  unidb.ErrUniqueViolation,
		int32(4):  unidb.ErrNoResult,
		int32(6):  unidb.ErrConnection,
		int32(7):  unidb.ErrConnection,
		int32(11): unidb.ErrAuthorizationFailed,
		int32(12): unidb.ErrInvalidSyntax, // NOTE: unsupported format
		int32(13): unidb.ErrAuthorizationFailed,
		int32(14): unidb.ErrInvalidSchemaName,
		int32(18): unidb.ErrAuthorizationFailed,
		int32(20): unidb.ErrCardinalityViolation,
		int32(21): unidb.ErrDataException,
		int32(22): unidb.ErrInvalidSyntax,
		int32(23): unidb.ErrTxState,
		int32(24): unidb.ErrTxState,
		int32(25): unidb.ErrRestrictViolation, // NOTE: Check what exactly remote validation error means
		int32(26): unidb.ErrDataException,
		int32(27): unidb.ErrDataException,
		int32(28): unidb.ErrInvalidSyntax,
		int32(29): unidb.ErrInvalidSyntax,
		int32(30): unidb.ErrInvalidSyntax,
		int32(31): unidb.ErrAuthorizationFailed,
		int32(32): unidb.ErrAuthorizationFailed,
		int32(33): unidb.ErrAuthorizationFailed,

		int32(39): unidb.ErrSystemError,
		int32(41): unidb.ErrSystemError,

		// NonMatchingDocument
		int32(47): unidb.ErrCheckViolation,
		int32(48): unidb.ErrInvalidSchemaName,
		int32(49): unidb.ErrAuthorizationFailed,
		int32(50): unidb.ErrTimedOut,
		int32(53): unidb.ErrDataException,
		int32(56): unidb.ErrInvalidResourceName,
		int32(57): unidb.ErrInvalidResourceName,
		int32(59): unidb.ErrInvalidSyntax,
		int32(66): unidb.ErrIntegrityConstraintViolation,
		int32(67): unidb.ErrUniqueViolation,
		int32(68): unidb.ErrUniqueViolation,

		int32(70): unidb.ErrReplica,
		int32(71): unidb.ErrReplica,
		int32(72): unidb.ErrCheckViolation,
		int32(73): unidb.ErrInvalidCatalogName,

		int32(74):  unidb.ErrReplica,
		int32(76):  unidb.ErrReplica,
		int32(78):  unidb.ErrRestrictViolation,
		int32(80):  unidb.ErrIntegrityConstraintViolation,
		int32(84):  unidb.ErrUniqueViolation,
		int32(85):  unidb.ErrIntegrityConstraintViolation,
		int32(86):  unidb.ErrIntegrityConstraintViolation,
		int32(89):  unidb.ErrConnection,
		int32(91):  unidb.ErrShutdown,
		int32(93):  unidb.ErrReplica,
		int32(98):  unidb.ErrSystemError,
		int32(103): unidb.ErrReplica,
		int32(107): unidb.ErrTxState,
		int32(111): unidb.ErrCheckViolation,
		int32(123): unidb.ErrReplica,
		int32(128): unidb.ErrTxState,
		int32(129): unidb.ErrTxState,
		// SSL
		int32(140): unidb.ErrConnection,
		int32(141): unidb.ErrConnection,
		int32(162): unidb.ErrNoResult,
		// Transport or transaction?
		int32(172): unidb.ErrConnection,
		int32(173): unidb.ErrConnection,
		int32(174): unidb.ErrConnection,

		int32(201):   unidb.ErrInternalError,
		int32(202):   unidb.ErrTimedOut,
		int32(207):   unidb.ErrCheckViolation,
		int32(211):   unidb.ErrNoResult,
		int32(212):   unidb.ErrTxRollback,
		int32(225):   unidb.ErrTxTermination,
		int32(251):   unidb.ErrTxNotFound,
		int32(256):   unidb.ErrTxDone,
		int32(257):   unidb.ErrTxState,
		int32(261):   unidb.ErrProgramLimitExceeded,
		int32(262):   unidb.ErrTimedOut,
		int32(264):   unidb.ErrSystemError,
		int32(267):   unidb.ErrTxBeginInProgress,
		int32(268):   unidb.ErrSystemError,
		int32(278):   unidb.ErrTxTermination,
		int32(279):   unidb.ErrConnection,
		int32(281):   unidb.ErrTxTermination,
		int32(282):   unidb.ErrTxTermination,
		int32(287):   unidb.ErrTxTermination,
		int32(11000): unidb.ErrUniqueViolation,
		int32(11600): unidb.ErrShutdown,
		int32(14031): unidb.ErrSystemError,
		int32(17280): unidb.ErrInvalidResourceName,
	}
}

/**


error_code("InternalError", 1)
error_code("BadValue", 2)
error_code("OBSOLETE_DuplicateKey", 3)
error_code("NoSuchKey", 4)
error_code("GraphContainsCycle", 5)
error_code("HostUnreachable", 6)
error_code("HostNotFound", 7)
error_code("UnknownError", 8)
error_code("FailedToParse", 9)
error_code("CannotMutateObject", 10)
error_code("UserNotFound", 11)
error_code("UnsupportedFormat", 12)
error_code("Unauthorized", 13)
error_code("TypeMismatch", 14)
error_code("Overflow", 15)
error_code("InvalidLength", 16)
*/
