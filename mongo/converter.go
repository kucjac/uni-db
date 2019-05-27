package mongo

import (
	"github.com/neuronlabs/uni-db"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

var _ unidb.Converter = &ErrorConverter{}

// ErrorConverter is the uni-db.Converter implementation of the mongo errors
type ErrorConverter struct {
	// mapping from errors to the unidb prototypes
	mapping map[interface{}]unidb.Error
}

// New creates new Error converter for the mongo db
func New() *ErrorConverter {
	return &ErrorConverter{mapping: defaultMapping()}
}

// Convert converts the mongo error into a *unidb.Error instance
func (e *ErrorConverter) Convert(err error) *unidb.Error {

	var outError *unidb.Error

	switch et := err.(type) {
	case *mongo.CommandError:
		pr, ok := e.mapping[et.Code]
		if ok {
			outError = pr.NewWithMessage(et.Message)
		}
	case *mongo.MarshalError:
		outError = unidb.ErrInvalidSyntax.NewWithMessage(et.Error())
	case *mongo.BulkWriteError:
		pr, ok := e.mapping[int32(et.Code)]
		if ok {
			outError = pr.NewWithMessage(et.Message)
		}
	case *mongo.WriteError:
		pr, ok := e.mapping[int32(et.Code)]
		if ok {
			outError = pr.NewWithMessage(et.Message)
		}
	case mongo.WriteErrors:
		var messages []string

		var hasUnique, hasSyntax bool
		for _, single := range et {
			messages = append(messages, single.Message)
			pr, ok := e.mapping[int32(single.Code)]
			if ok {
				if pr.ID == unidb.ErrUniqueViolation.ID {
					hasUnique = true
				} else if pr.ID == unidb.ErrInvalidSyntax.ID {
					hasSyntax = true
				}
			}
		}

		switch {
		case hasUnique:
			outError = unidb.ErrUniqueViolation.NewWithMessage(strings.Join(messages, ","))
		case hasSyntax:
			outError = unidb.ErrInvalidSyntax.NewWithMessage(strings.Join(messages, ","))
		default:
			outError = unidb.ErrCheckViolation.NewWithMessage(strings.Join(messages, ","))
		}

	default:
		pr, ok := e.mapping[err]
		if ok {
			outError = pr.NewWithError(err)
		} else {
			outError = unidb.ErrUnspecifiedError.NewWithError(err)
		}
	}
	return outError
}

func (e *ErrorConverter) setMapping(mapping map[interface{}]unidb.Error) {
	e.mapping = mapping
}
