package gormock

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/xorcare/gormock/comp"
)

// Gormock implements the interface Foundation and describes
// the behavior of the handlers used by default in gormock.
type Gormock struct {
	mock.Mock
}

// New returns new instance of Gormock.
func New(t mock.TestingT) (g *Gormock) {
	g = &Gormock{}
	g.Test(t)
	return g
}

// DB returns an instance db is compatible on the gorm.db.
func (g *Gormock) DB() *comp.DB {
	return &comp.DB{
		Foundation: g,
	}
}

// Scope returns an instance Scope is compatible on the gorm.Scope.
func (g *Gormock) Scope() *comp.Scope {
	return &comp.Scope{
		Foundation: g,
	}
}

// Handler handles calls that do not require return values.
// Implementation of Foundation interface.
func (g *Gormock) Handler(
	receiver interface{},
	methodName string,
	arguments ...interface{},
) {
	g.HandlerOther(receiver, methodName, nil, arguments)
}

// HandlerOther handles calls that return a single result value.
// Implementation of Foundation interface.
func (g *Gormock) HandlerOther(
	receiver interface{},
	methodName string,
	outcome interface{},
	arguments ...interface{},
) interface{} {
	type resolver = func(receiver interface{}) (result interface{})
	type resolverArgs = func(
		receiver interface{},
		arguments ...interface{},
	) (result interface{})

	_, isDB := outcome.(*comp.DB)
	if receiver, ok := receiver.(*comp.DB); ok && isDB {
		return g.handlerDB(receiver, methodName, arguments...)
	}

	_, isScope := outcome.(*comp.Scope)
	if receiver, ok := receiver.(*comp.DB); ok && isScope {
		return g.newScope(receiver, methodName, arguments...)
	}

	rvs := g.MethodCalled(methodName, arguments...)
	if len(rvs) > 0 {
		reply := rvs[0]
		switch reply.(type) {
		case resolver:
			return reply.(resolver)(receiver)
		case resolverArgs:
			return reply.(resolverArgs)(receiver, arguments...)
		default:
			return reply
		}
	}

	return outcome
}

// HandlerOthers handles calls that return two result values.
// Implementation of Foundation interface.
func (g *Gormock) HandlerOthers(
	receiver interface{},
	methodName string,
	outcomeFirst, outcomeSecond interface{},
	arguments ...interface{},
) (interface{}, interface{}) {
	type resolver = func(
		receiver interface{},
	) (resultFirst, resultSecond interface{})
	type resolverArgs = func(
		receiver interface{},
		arguments ...interface{},
	) (resultFirst, resultSecond interface{})

	rvs := g.MethodCalled(methodName, arguments...)
	if len(rvs) > 0 {
		reply := rvs[0]
		switch reply.(type) {
		case resolver:
			return reply.(resolver)(receiver)
		case resolverArgs:
			return reply.(resolverArgs)(receiver, arguments...)
		default:
			return rvs[0], rvs[1]
		}
	}

	return outcomeFirst, outcomeSecond
}

func (g *Gormock) handlerDB(
	receiver *comp.DB,
	methodName string,
	arguments ...interface{},
) (outcome *comp.DB) {
	type resolver = func(receiver *comp.DB) (result *comp.DB)
	type resolverArgs = func(
		receiver *comp.DB,
		arguments ...interface{},
	) (result *comp.DB)

	outcome = receiver
	rvs := g.MethodCalled(methodName, arguments...)
	if len(rvs) > 0 {
		reply := rvs[0]
		switch reply.(type) {
		case comp.DB:
			db := reply.(comp.DB)
			outcome = &db
		case *comp.DB:
			outcome = reply.(*comp.DB)
		case nil:
			outcome = nil
		case resolver:
			return reply.(resolver)(receiver)
		case resolverArgs:
			return reply.(resolverArgs)(receiver, arguments...)
		default:
			panic(fmt.Sprintf("call() unknown type %T", reply))
		}
	}

	return outcome
}

func (g *Gormock) newScope(
	receiver *comp.DB,
	methodName string,
	arguments ...interface{},
) (outcome *comp.Scope) {
	type resolver = func(receiver *comp.DB) (result *comp.Scope)
	type resolverArgs = func(
		receiver *comp.DB,
		arguments ...interface{},
	) (result *comp.Scope)

	outcome = &comp.Scope{
		Foundation: receiver.Foundation,
		PDB:        receiver,
	}
	rvs := g.MethodCalled(methodName, arguments...)
	if len(rvs) > 0 {
		reply := rvs[0]
		switch reply.(type) {
		case comp.DB:
			scope := reply.(comp.Scope)
			if scope.PDB == nil {
				scope.PDB = receiver
			}
			outcome = &scope
		case *comp.Scope:
			scope := reply.(*comp.Scope)
			scope.PDB = receiver
			if scope != nil && scope.PDB == nil {
				scope.PDB = receiver
			}
			outcome = scope
		case nil:
			outcome = nil
		case resolver:
			return reply.(resolver)(receiver)
		case resolverArgs:
			return reply.(resolverArgs)(receiver, arguments...)
		default:
			panic(fmt.Sprintf("unknown type %T", reply))
		}
	}

	return outcome
}
