package unidb

// Converter is an interface that defines the form
// of converting third-party database errors into uniform
// dberrors.Error. The returned errors should be based on the
// prototypes provided in this package.
type Converter interface {
	Convert(err error) *Error
}
