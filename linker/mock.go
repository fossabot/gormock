// +build gormock

package linker

import (
	"github.com/xorcare/gormock"
)

// DB link to mock.
type DB = gormock.DB

// Association type alias.
type Association = gormock.Association

// Callback type alias.
type Callback = gormock.Callback

// Dialect type alias.
type Dialect = gormock.Dialect

// Field how to handle many2many relations.
type Field = gormock.Field

// JoinTableHandlerInterface type alias.
type JoinTableHandlerInterface = gormock.JoinTableHandlerInterface

// Scope type alias.
type Scope = gormock.Scope

// SQLCommon type alias.
type SQLCommon = gormock.SQLCommon
