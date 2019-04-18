package pgconv

import (
	"database/sql"
	"fmt"
	"github.com/kucjac/uni-db"
	"github.com/lib/pq"
)

// PGConverter is an implementation of dberrorsrrors.Converter.
type PGConverter struct {
	errorMap map[interface{}]unidb.Error
}

// Convert converts the given error into *Error.
// The method checks if given error is of known type, and then returns it.ty
// If an error is unknown it returns new 'dberrorsrrors.ErrUnspecifiedError'.
// At first converter checks if an error is of *pq.Error type.
// Having a postgres *pq.Error it checks if an ErrorCode is in the map,
// and returns it if true. Otherwise method checks if the ErrorClass exists in map.
// If it is present, new *Error of given type is returned.
func (p *PGConverter) Convert(err error) (dberrorsErr *unidb.Error) {
	pgError, ok := err.(*pq.Error)
	if !ok {
		// The error may be of sql.ErrNoRows type
		if err == sql.ErrNoRows {
			return unidb.ErrNoResult.NewWithError(err)
		} else if err == sql.ErrTxDone {
			return unidb.ErrTxDone.NewWithError(err)
		}
		return unidb.ErrUnspecifiedError.NewWithError(err)

	}

	// Error prototype
	var dbErrorProto unidb.Error

	fmt.Printf("PgError: %#v\n", pgError)
	fmt.Printf("Converter: %#v\n", p)

	// First check if recogniser has entire error code in it
	dbErrorProto, ok = p.errorMap[pgError.Code]
	if ok {
		return dbErrorProto.NewWithError(err)
	}

	// If the ErrorCode is not present, check the code class
	dbErrorProto, ok = p.errorMap[pgError.Code.Class()]
	if ok {
		return dbErrorProto.NewWithError(err)
	}

	// If the Error Class is not presen in the error map
	// return ErrDBNotMapped
	return unidb.ErrUnspecifiedError.NewWithError(err)
}

// New creates new PGConverter
// It is already inited and ready to use.
func New() *PGConverter {
	return &PGConverter{errorMap: defaultPGErrorMap}
}
