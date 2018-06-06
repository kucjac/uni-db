package gormconv

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/kucjac/uni-db"
	"github.com/kucjac/uni-db/mysqlconv"
	"github.com/kucjac/uni-db/pgconv"
)

// GORMConverter defines error converter for multiple databases drivers
// used by the 'gorm' package.
// Implements 'Converter' interface.
type GORMConverter struct {
	converter unidb.Converter
}

// New creates new *GORMConverter.
// On the base of the *gorm.DB argument it recognise given gorm.Dialect
// on the base of the dialect the function recognise the appropiate error converter.
// returns error if the nil pointer provided or unsupported db.Dialect
func New(db *gorm.DB) (conv *GORMConverter, err error) {
	conv = &GORMConverter{}
	err = conv.initialize(db)
	if err != nil {
		return nil, err
	}
	return conv, nil
}

// Convert implements unidb.Converter
// Converts provided argument error into *unidb.Error type
func (g *GORMConverter) Convert(err error) (dbErr *unidb.Error) {
	switch err {
	case gorm.ErrCantStartTransaction, gorm.ErrInvalidTransaction:
		dbErr = unidb.ErrInvalidTransState.NewWithError(err)
	case gorm.ErrInvalidSQL:
		dbErr = unidb.ErrInvalidSyntax.NewWithError(err)
	case gorm.ErrUnaddressable:
		dbErr = unidb.ErrUnspecifiedError.NewWithError(err)
	case gorm.ErrRecordNotFound:
		dbErr = unidb.ErrNoResult.NewWithError(err)
	}
	if dbErr == nil {
		dbErr = g.converter.Convert(err)
	}
	dbErr.Message = err.Error()
	// If error is not of gorm type
	// use db recogniser
	return dbErr
}

// initialize provides initialization process of the GORMConverter
// returns error if nil pointer provided or unsupported database dialect.
func (g *GORMConverter) initialize(db *gorm.DB) error {
	if db == nil {
		return errors.New("Nil pointer provided")
	}
	dialect := db.Dialect()
	switch dialect.GetName() {
	case "postgres":
		g.converter = pgconv.New()
	case "mysql":
		g.converter = mysqlconv.New()
	case "sqlite3":
		g.converter = AnyConverter{}
	default:
		return errors.New("Unsupported database dialect.")
	}
	return nil
}
