package gormock

import (
	"fmt"
)

// Kernel implements the interface Foundation and describes
// the behavior of the handlers used by default in gormock.
type Kernel struct{}

// Handler handles calls that do not require return values.
func (k Kernel) Handler(
	receiver Receiver,
	methodName string,
	arguments ...interface{},
) {
	k.HandlerOther(receiver, methodName, nil, arguments)
}

// HandlerOther handles calls that return a single result value.
func (k Kernel) HandlerOther(
	receiver Receiver,
	methodName string,
	outcome interface{},
	arguments ...interface{},
) interface{} {
	type resolver = func(receiver interface{}) (result interface{})
	type resolverArgs = func(
		receiver interface{},
		arguments ...interface{},
	) (result interface{})

	_, isDB := outcome.(*DB)
	if receiver, ok := receiver.(*DB); ok && isDB {
		return k.handlerDB(receiver, methodName, arguments...)
	}

	_, isScope := outcome.(*Scope)
	if receiver, ok := receiver.(*DB); ok && isScope {
		return k.newScope(receiver, methodName, arguments...)
	}

	rvs := receiver.MethodCalled(methodName, arguments...)
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
func (Kernel) HandlerOthers(
	receiver Receiver,
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

	rvs := receiver.MethodCalled(methodName, arguments...)
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

func (Kernel) handlerDB(
	receiver *DB,
	methodName string,
	arguments ...interface{},
) (outcome *DB) {
	type resolver = func(receiver *DB) (result *DB)
	type resolverArgs = func(
		receiver *DB,
		arguments ...interface{},
	) (result *DB)

	outcome = receiver
	rvs := receiver.MethodCalled(methodName, arguments...)
	if len(rvs) > 0 {
		reply := rvs[0]
		switch reply.(type) {
		case DB:
			db := reply.(DB)
			db.Parent = receiver
			outcome = &db
		case *DB:
			db := reply.(*DB)
			db.Parent = receiver
			outcome = db
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
func (Kernel) newScope(
	receiver *DB,
	methodName string,
	arguments ...interface{},
) (outcome *Scope) {
	type resolver = func(receiver *DB) (result *Scope)
	type resolverArgs = func(
		receiver *DB,
		arguments ...interface{},
	) (result *Scope)

	outcome = &Scope{
		Foundation: receiver.Foundation,
		Mock:       receiver.Mock,
		db:         receiver,
	}
	rvs := receiver.MethodCalled(methodName, arguments...)
	if len(rvs) > 0 {
		reply := rvs[0]
		switch reply.(type) {
		case DB:
			scope := reply.(Scope)
			scope.db = receiver
			outcome = &scope
		case *Scope:
			scope := reply.(*Scope)
			scope.db = receiver
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
