package sqliteconv

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"github.com/neuronlabs/uni-db"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewConverter(t *testing.T) {
	Convey("Using New function creates non-empty SQLiteConverter.", t, func() {
		var converter *SQLiteConverter
		converter = New()

		So(converter, ShouldNotBeNil)

		So(len(converter.errorMap), ShouldBeGreaterThan, 0)

		Convey("The SQLiteConverter implements Converter", func() {
			So(converter, ShouldImplement, (*unidb.Converter)(nil))
		})
	})
}

func TestSQLiteRecogniser(t *testing.T) {
	Convey("Using a SQLite3 Error Converter", t, func() {

		var converter *SQLiteConverter = New()

		Convey("Having a list of sqlite errors", func() {

			sqliteErrors := map[sqlite3.Error]unidb.Error{
				{Code: sqlite3.ErrWarning}:  unidb.ErrWarning,
				{Code: sqlite3.ErrNotFound}: unidb.ErrNoResult,
				{Code: sqlite3.ErrCantOpen}: unidb.ErrConnection,
				{Code: sqlite3.ErrNotADB}:   unidb.ErrConnection,
				{Code: sqlite3.ErrMismatch}: unidb.ErrDataException,
				{Code: sqlite3.ErrConstraint,
					ExtendedCode: sqlite3.ErrConstraintPrimaryKey}: unidb.ErrUniqueViolation,
				{Code: sqlite3.ErrConstraint,
					ExtendedCode: sqlite3.ErrConstraintFunction}: unidb.ErrIntegrityConstraintViolation,
				{Code: sqlite3.ErrConstraint,
					ExtendedCode: sqlite3.ErrConstraintCheck}: unidb.ErrCheckViolation,
				{Code: sqlite3.ErrConstraint,
					ExtendedCode: sqlite3.ErrConstraintForeignKey}: unidb.ErrForeignKeyViolation,
				{Code: sqlite3.ErrConstraint,
					ExtendedCode: sqlite3.ErrConstraintUnique}: unidb.ErrUniqueViolation,
				{Code: sqlite3.ErrConstraint,
					ExtendedCode: sqlite3.ErrConstraintNotNull}: unidb.ErrNotNullViolation,
				{Code: sqlite3.ErrProtocol}: unidb.ErrTxState,
				{Code: sqlite3.ErrRange}:    unidb.ErrInvalidSyntax,
				{Code: sqlite3.ErrError}:    unidb.ErrInvalidSyntax,
				{Code: sqlite3.ErrAuth}:     unidb.ErrAuthorizationFailed,
				{Code: sqlite3.ErrPerm}:     unidb.ErrInsufficientPrivilege,
				{Code: sqlite3.ErrFull}:     unidb.ErrInsufficientResources,
				{Code: sqlite3.ErrTooBig}:   unidb.ErrProgramLimitExceeded,
				{Code: sqlite3.ErrNoLFS}:    unidb.ErrSystemError,
				{Code: sqlite3.ErrInternal}: unidb.ErrInternalError,
				{Code: sqlite3.ErrEmpty}:    unidb.ErrUnspecifiedError,
			}

			Convey("For given *sqlite.Error, specific database error should be returner.", func() {
				for sqliteErr, dbErr := range sqliteErrors {
					recognisedErr := converter.Convert(sqliteErr)
					Println(sqliteErr)
					Println(recognisedErr)
					Println(dbErr)
					So(recognisedErr.Compare(dbErr), ShouldBeTrue)
				}
			})
		})

		Convey("Having an error of type sql.Err*, error is converted into *Error type.", func() {
			var err error
			err = sql.ErrNoRows
			recognisedErr := converter.Convert(err)
			So(recognisedErr.Compare(unidb.ErrNoResult), ShouldBeTrue)

			err = sql.ErrTxDone
			recognisedErr = converter.Convert(err)
			So(recognisedErr.Compare(unidb.ErrTxDone), ShouldBeTrue)
		})

		Convey("Having an error of different type than *sqlite3.Error and sql.Err*", func() {
			err := errors.New("Unknown error type")
			recognisedErr := converter.Convert(err)
			So(recognisedErr.Compare(unidb.ErrUnspecifiedError), ShouldBeTrue)
		})
	})
}
