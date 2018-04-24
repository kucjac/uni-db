package mysqlconv

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/kucjac/uni-db"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNew(t *testing.T) {
	Convey("The 'New()' function creates new already inited '*MySQLConverter' entity", t, func() {
		var converter *MySQLConverter
		converter = New()

		So(converter, ShouldNotBeNil)

		So(len(converter.codeMap), ShouldBeGreaterThan, 0)
		So(len(converter.sqlStateMap), ShouldBeGreaterThan, 0)

		Convey("The *MySQLConverter implements unidb.Converter interface", func() {
			So(converter, ShouldImplement, (*unidb.Converter)(nil))
		})
	})
}

func TestMySQLRecogniser(t *testing.T) {
	Convey("Having MySQLConverter.", t, func() {
		var converter *MySQLConverter = New()

		Convey("Check if selected MySQL Errors would return for given unidb Error", func() {
			errorMap := map[*mysql.MySQLError]unidb.Error{
				{Number: 1022}: unidb.ErrUniqueViolation,
				{Number: 1046}: unidb.ErrInvalidCatalogName,
				{Number: 1048}: unidb.ErrNotNullViolation,
				{Number: 1050}: unidb.ErrInvalidSyntax,
				{Number: 1062}: unidb.ErrUniqueViolation,
				{Number: 1114}: unidb.ErrProgramLimitExceeded,
				{Number: 1118}: unidb.ErrProgramLimitExceeded,
				{Number: 1129}: unidb.ErrInternalError,
				{Number: 1130}: unidb.ErrInvalidAuthorization,
				{Number: 1131}: unidb.ErrInvalidAuthorization,
				{Number: 1132}: unidb.ErrInvalidPassword,
				{Number: 1133}: unidb.ErrInvalidPassword,
				{Number: 1169}: unidb.ErrUniqueViolation,
				{Number: 1182}: unidb.ErrTransRollback,
				{Number: 1216}: unidb.ErrForeignKeyViolation,
				{Number: 1217}: unidb.ErrForeignKeyViolation,
				{Number: 1227}: unidb.ErrInsufficientPrivilege,
				{Number: 1251}: unidb.ErrInvalidAuthorization,
				{Number: 1400}: unidb.ErrInvalidTransState,
				{Number: 1401}: unidb.ErrInternalError,
				{Number: 1451}: unidb.ErrForeignKeyViolation,
				{Number: 1452}: unidb.ErrForeignKeyViolation,
				{Number: 1557}: unidb.ErrUniqueViolation,
				{Number: 1568}: unidb.ErrUniqueViolation,
				{Number: 1698}: unidb.ErrInvalidPassword,
				//Nil
				{Number: 1317}: unidb.ErrUnspecifiedError,
				{Number: 1040}: unidb.ErrConnExc,
				//Non mapped number
				{Number: 1000}: unidb.ErrUnspecifiedError,
			}

			for msqlErr, dbErr := range errorMap {
				dbErrInMap := converter.Convert(msqlErr)

				So(dbErrInMap.Compare(dbErr), ShouldBeTrue)
			}
		})
		Convey("Having error of different type than *mysql.Error", func() {
			errorMap := map[error]unidb.Error{
				sql.ErrNoRows:           unidb.ErrNoResult,
				sql.ErrTxDone:           unidb.ErrTxDone,
				mysql.ErrInvalidConn:    unidb.ErrConnExc,
				mysql.ErrNoTLS:          unidb.ErrConnExc,
				mysql.ErrMalformPkt:     unidb.ErrConnExc,
				mysql.ErrOldProtocol:    unidb.ErrConnExc,
				mysql.ErrNativePassword: unidb.ErrUnspecifiedError,
			}

			for err, dbErr := range errorMap {
				dbErrInMap := converter.Convert(err)
				// Printf("%v: %v\n", err, dbErrInMap)
				So(dbErrInMap.Compare(dbErr), ShouldBeTrue)
			}
		})
	})
}
