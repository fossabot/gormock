// Copyright © 2018-2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gormock

package linker

import (
	"github.com/jinzhu/gorm"
)

// DB type alias.
type DB = gorm.DB

// Association type alias.
type Association = gorm.Association

// Callback type alias.
type Callback = gorm.Callback

// CallbackProcessor type alias.
type CallbackProcessor = gorm.CallbackProcessor

// Dialect type alias.
type Dialect = gorm.Dialect

// Field type alias.
type Field = gorm.Field

// JoinTableHandlerInterface type alias.
type JoinTableHandlerInterface = gorm.JoinTableHandlerInterface

// Scope type alias.
type Scope = gorm.Scope

// SQLCommon type alias.
type SQLCommon = gorm.SQLCommon
