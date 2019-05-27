package gormconv

import (
	"github.com/neuronlabs/uni-db"
)

type AnyConverter struct{}

func (a AnyConverter) Convert(err error) *unidb.Error {
	return unidb.ErrUnspecifiedError.New()
}
