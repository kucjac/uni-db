package unidb

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

// errorPrototypes all possible error prototypes
var errorPrototypes []Error = []Error{
	ErrWarning,
	ErrNoResult,
	ErrConnExc,
	ErrCardinalityViolation,
	ErrDataException,
	ErrIntegrConstViolation,
	ErrRestrictViolation,
	ErrNotNullViolation,
	ErrForeignKeyViolation,
	ErrUniqueViolation,
	ErrCheckViolation,
	ErrInvalidTransState,
	ErrInvalidTransTerm,
	ErrTransRollback,
	ErrTxDone,
	ErrInvalidAuthorization,
	ErrInvalidPassword,
	ErrInvalidSchemaName,
	ErrInvalidCatalogName,
	ErrInvalidSyntax,
	ErrInsufficientPrivilege,
	ErrInsufficientResources,
	ErrProgramLimitExceeded,
	ErrSystemError,
	ErrInternalError,
	ErrUnspecifiedError,
}

func TestNew(t *testing.T) {
	Convey("Having Error prototypes", t, func() {
		var randomMessage string
		var randomError error
		var createdErr *Error

		Convey("Using New method on prototype should create a new *Error entity with the same id and title as prototype", func() {
			for _, proto := range errorPrototypes {
				createdErr = proto.New()
				So(createdErr.ID, ShouldEqual, proto.ID)
				So(createdErr.Title, ShouldEqual, proto.Title)
				So(createdErr, ShouldNotEqual, &proto)
			}
		})

		Convey("Using NewWithMessage method should create a new *Error with the same id, title as prototype with the Message provided in the argument", func() {
			for _, proto := range errorPrototypes {
				randomMessage = newRandomMessage()
				createdErr = proto.NewWithMessage(randomMessage)

				So(createdErr.ID, ShouldEqual, proto.ID)
				So(createdErr.Title, ShouldEqual, proto.Title)
				So(createdErr.Message, ShouldEqual, randomMessage)
				So(createdErr, ShouldNotEqual, &proto)
			}

		})

		Convey("Using NewWithError method should create a new *Error with the same id, title. Provided 'err' argument should be saved in Message field.", func() {
			for _, proto := range errorPrototypes {
				randomError = newRandomError()

				createdErr = proto.NewWithError(randomError)

				So(createdErr.ID, ShouldEqual, proto.ID)
				So(createdErr.Title, ShouldEqual, proto.Title)
				So(createdErr.Message, ShouldEqual, randomError.Error())
				So(createdErr, ShouldNotEqual, &proto)
			}
		})

	})
}

func TestCompare(t *testing.T) {
	Convey("Having error prototypes", t, func() {
		var ok bool
		var dbErr *Error
		Convey(`If *Error has the same ID as Error prototype 
		the Compare method should return true and`, func() {

			for _, proto := range errorPrototypes {
				dbErr = &Error{ID: proto.ID}

				ok = dbErr.Compare(proto)

				So(ok, ShouldBeTrue)
			}

		})
		Convey("Otherwise it should return false", func() {
			for i := 0; i <= len(errorPrototypes)-2; i++ {
				dbErr = &Error{ID: errorPrototypes[i].ID}

				ok = dbErr.Compare(errorPrototypes[i+1])

				So(ok, ShouldBeFalse)
			}
		})
	})

}

func TestGetPrototype(t *testing.T) {
	Convey(`Having the error entity with the same id as given prototype, 
		GetPrototype returns the proto`, t, func() {
		for _, proto := range errorPrototypes {
			errorEntity := &Error{ID: proto.ID}
			retrievedProto, err := errorEntity.GetPrototype()

			So(err, ShouldBeNil)
			So(retrievedProto, ShouldResemble, proto)
		}
	})

	Convey(`If the Error entity contains unrecognisable ID, 
	GetPrototype returns error`, t, func() {
		dbError := &Error{Title: "There is no ID field"}

		_, err := dbError.GetPrototype()
		So(err, ShouldBeError)

		dbError = &Error{ID: 131224}

		_, err = dbError.GetPrototype()
		So(err, ShouldBeError)

	})
}

func newRandomError() error {
	return errors.New(randSeq(20))
}

func newRandomMessage() string {
	return randSeq(20)
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters)-1)]
	}

	return string(b)
}
