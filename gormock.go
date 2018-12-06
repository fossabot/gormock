package gormock

import (
	"github.com/stretchr/testify/mock"
)

// Basic a global variable that provides the ability to replace the
// implementation of method invocation handlers is global for all instances
// created, this allows you to make packages that can replace the
// implementation through the init function using a unnamed import.
var Basic Foundation = Kernel{}

// Receiver an interface that will be used primarily by mock and its children
// to generalize method calls from the kernel implementation.
type Receiver interface {
	MethodCalled(string, ...interface{}) mock.Arguments
}

// Foundation an interface describing methods that implement all the basic
// processing of called methods.
type Foundation interface {
	// Handler handles calls that do not require return values.
	Handler(
		receiver Receiver,
		methodName string,
		arguments ...interface{},
	)
	// HandlerOther handles calls that return a single result value.
	HandlerOther(
		receiver Receiver,
		methodName string,
		outcome interface{},
		arguments ...interface{},
	) (result interface{})
	// HandlerOthers handles calls that return two result values.
	HandlerOthers(
		receiver Receiver,
		methodName string,
		outcomeFirst, outcomeSecond interface{},
		arguments ...interface{},
	) (resultFirst, resultSecond interface{})
}

// DB a type that combines the capabilities of a mock and an interface
// that is almost fully compatible with gorm.
type DB struct {
	// Mock fields definition
	Foundation
	*mock.Mock
	Parent *DB

	// Gorm fields definition
	Value        interface{}
	Error        error
	RowsAffected int64
}

// New a constructor that provides a new customized instance of a mock.
func New(t mock.TestingT) *DB {
	m := &mock.Mock{}
	m.Test(t)

	return &DB{
		Mock:       m,
		Foundation: Basic,
	}
}

// clone the method for cloning an object, which makes a copy of the object
// and sets a pointer to the parent object, is designed to preserve some of
// the behavior of the original object and reduce problems with race condition.
func (s *DB) clone() (out *DB) {
	out = &DB{}
	*out = *s
	out.Parent = s
	return
}
