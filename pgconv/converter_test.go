package pgconv

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/neuronlabs/uni-db"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewConverter(t *testing.T) {
	Convey("While using New function the converter is already inited.", t, func() {
		var converter *PGConverter
		converter = New()
		So(converter, ShouldNotBeNil)
		So(len(converter.errorMap), ShouldBeGreaterThan, 0)
		So(converter, ShouldImplement, (*unidb.Converter)(nil))
	})
}

func TestPGRecogniser(t *testing.T) {
	Convey("Using Postgress Converter", t, func() {
		var converter *PGConverter = New()

		Convey("Having a list of typical postgres errors", func() {
			postgresErrors := map[*pq.Error]unidb.Error{
				{Code: pq.ErrorCode("01000")}: unidb.ErrWarning,
				{Code: pq.ErrorCode("01007")}: unidb.ErrWarning,
				{Code: pq.ErrorCode("02000")}: unidb.ErrNoResult,
				{Code: pq.ErrorCode("P0002")}: unidb.ErrNoResult,
				{Code: pq.ErrorCode("08006")}: unidb.ErrConnection,
				{Code: pq.ErrorCode("21000")}: unidb.ErrCardinalityViolation,
				{Code: pq.ErrorCode("22012")}: unidb.ErrDataException,
				{Code: pq.ErrorCode("23000")}: unidb.ErrIntegrityConstraintViolation,
				{Code: pq.ErrorCode("23001")}: unidb.ErrRestrictViolation,
				{Code: pq.ErrorCode("23502")}: unidb.ErrNotNullViolation,
				{Code: pq.ErrorCode("23503")}: unidb.ErrForeignKeyViolation,
				{Code: pq.ErrorCode("23505")}: unidb.ErrUniqueViolation,
				{Code: pq.ErrorCode("23514")}: unidb.ErrCheckViolation,
				{Code: pq.ErrorCode("25001")}: unidb.ErrTxState,
				{Code: pq.ErrorCode("25004")}: unidb.ErrTxState,
				{Code: pq.ErrorCode("28000")}: unidb.ErrAuthorizationFailed,
				{Code: pq.ErrorCode("28P01")}: unidb.ErrAuthenticationFailed,
				{Code: pq.ErrorCode("2D000")}: unidb.ErrTxTermination,
				{Code: pq.ErrorCode("3F000")}: unidb.ErrInvalidSchemaName,
				{Code: pq.ErrorCode("40000")}: unidb.ErrTxRollback,
				{Code: pq.ErrorCode("42P06")}: unidb.ErrInvalidSyntax,
				{Code: pq.ErrorCode("42501")}: unidb.ErrInsufficientPrivilege,
				{Code: pq.ErrorCode("53100")}: unidb.ErrInsufficientResources,
				{Code: pq.ErrorCode("54011")}: unidb.ErrProgramLimitExceeded,
				{Code: pq.ErrorCode("58000")}: unidb.ErrSystemError,
				{Code: pq.ErrorCode("XX000")}: unidb.ErrInternalError,
				{Code: pq.ErrorCode("P0003")}: unidb.ErrUnspecifiedError,
			}

			Convey("For given postgres error, specific database error should return", func() {
				for pgErr, dbErr := range postgresErrors {
					convertedErr := converter.Convert(pgErr)
					So(convertedErr.Compare(dbErr), ShouldBeTrue)
				}
			})
		})

		Convey("When sql errors are returned, they are also converted into dberror", func() {
			errNoResults := converter.Convert(sql.ErrNoRows)
			So(errNoResults.Compare(unidb.ErrNoResult), ShouldBeTrue)

			errTxDone := converter.Convert(sql.ErrTxDone)
			So(errTxDone.Compare(unidb.ErrTxDone), ShouldBeTrue)
		})

		Convey("Having unknown error not of *pq.Error type forwards it", func() {

			fwdErr := converter.Convert(errors.New("Forwarded"))
			So(fwdErr.Compare(unidb.ErrUnspecifiedError), ShouldBeTrue)

		})
	})

}
