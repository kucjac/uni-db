package gormconv

import (
	"database/sql"
	"errors"
	"github.com/kucjac/uni-db"
	"github.com/kucjac/uni-db/gormconv"
	_ "github.com/kucjac/uni-db/gormconv/dialects/mysql"
	_ "github.com/kucjac/uni-db/gormconv/dialects/postgres"
	_ "github.com/kucjac/uni-db/gormconv/dialects/sqlite"
	"github.com/neuronlabs/gorm"
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestNewGormConverter(t *testing.T) {

	Convey("Subject: Creating new GORMConverter and initialize it with Init method", t, func() {

		Convey("Having *GORMConverter entity and some *gorm.DB connections", func() {
			var errorConverter *gormconv.GORMConverter

			dbSqlite, _ := gorm.Open("sqlite3", "./tests.db")
			dbPostgres, _ := gorm.Open("postgres", "host=myhost port=myport")
			dbMySQL, _ := gorm.Open("mysql", "user:password@/dbname")
			dbMSSQL, _ := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")

			gormSupported := []*gorm.DB{dbSqlite, dbPostgres, dbMySQL}

			Convey("While using Init method", func() {
				var err error

				Convey("If the dialect is supported, specific converter would be set", func() {
					for _, db := range gormSupported {
						errorConverter, err = gormconv.New(db.Dialect().GetName())
						So(err, ShouldBeNil)
						So(errorConverter, ShouldImplement, (*unidb.Converter)(nil))
					}
				})

				Convey("If the dialect is unsupported an error would be returned", func() {
					errorConverter, err = gormconv.New(dbMSSQL.Dialect().GetName())
					So(err, ShouldBeError)
					So(errorConverter, ShouldBeNil)
				})

			})

		})

	})

}

func TestGORMConverterConvert(t *testing.T) {

	Convey("Subject: Converting an error into *Error using GormErrorConverter method Convert", t, func() {

		Convey("Having inited GORMConverter", func() {
			db, err := gorm.Open("sqlite3", "./tests.db")
			So(err, ShouldBeNil)

			errorConverter, err := gormconv.New(db.Dialect().GetName())
			So(err, ShouldBeNil)

			Convey("Providing any error would result with *DBerror", func() {
				convertErrors := []error{gorm.ErrCantStartTransaction,
					gorm.ErrInvalidTransaction,
					gorm.ErrInvalidSQL,
					gorm.ErrUnaddressable,
					gorm.ErrRecordNotFound,
					unidb.ErrCardinalityViolation.New(),
					unidb.ErrWarning.New(),
					errors.New("Some error"),
					sql.ErrNoRows,
				}

				for _, err := range convertErrors {
					converted := errorConverter.Convert(err)
					So(converted, ShouldHaveSameTypeAs, &unidb.Error{})
				}
			})

		})

	})

}
