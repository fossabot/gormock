package comp

import (
	"github.com/jinzhu/gorm"
)

// Callback type alias.
type Callback = gorm.Callback

// Dialect type alias.
type Dialect = gorm.Dialect

// Field type alias.
type Field = gorm.Field

// JoinTableHandlerInterface type alias.
type JoinTableHandlerInterface = gorm.JoinTableHandlerInterface

// Search type alias.
type Search = interface{}

// SQLCommon type alias.
type SQLCommon = gorm.SQLCommon
