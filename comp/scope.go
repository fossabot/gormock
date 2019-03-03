// Copyright © 2018-2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comp

import (
	"reflect"
)

// Scope a type that combines the capabilities of a mock and an interface that
// is almost fully compatible with gorm.Scope.
type Scope struct {
	// Mock fields definition
	Foundation
	PDB *DB

	// Scope public fields
	Search  Search
	Value   interface{}
	SQL     string
	SQLVars []interface{}
}

// IndirectValue implementation of Scope interface.
func (r *Scope) IndirectValue() (value reflect.Value) {
	return r.HandlerOther(r, "IndirectValue", value).(reflect.Value)
}

// New implementation of Scope interface.
func (r *Scope) New(value interface{}) (o *Scope) {
	return r.HandlerOther(r, "New", o, value).(*Scope)
}

// DB returns the parent DB or nil if the current Scope or parent DB
// is not initialized. Implementation of Scope interface.
func (r *Scope) DB() *DB {
	if r == nil {
		return nil
	}

	return r.PDB
}

// NewDB implementation of Scope interface.
func (r *Scope) NewDB() (db *DB) {
	return r.HandlerOther(r, "NewDB", db).(*DB)
}

// SQLDB implementation of Scope interface.
func (r *Scope) SQLDB() (sqlCommon SQLCommon) {
	return r.HandlerOther(r, "SQLDB", sqlCommon).(SQLCommon)
}

// Dialect implementation of Scope interface.
func (r *Scope) Dialect() (dialect Dialect) {
	return r.HandlerOther(r, "Dialect", dialect).(Dialect)
}

// Quote implementation of Scope interface.
func (r *Scope) Quote(str string) (o string) {
	return r.HandlerOther(r, "Quote", o).(string)
}

// Err implementation of Scope interface.
func (r *Scope) Err(err error) (o error) {
	return r.HandlerOther(r, "Err", o, err).(error)
}

// HasError implementation of Scope interface.
func (r *Scope) HasError() (o bool) {
	return r.HandlerOther(r, "HasError", o).(bool)
}

// Log implementation of Scope interface.
func (r *Scope) Log(v ...interface{}) {
	r.Handler(r, "Log", v)
}

// SkipLeft implementation of Scope interface.
func (r *Scope) SkipLeft() {
	r.Handler(r, "SkipLeft")
}

// Fields implementation of Scope interface.
func (r *Scope) Fields() (fields []*Field) {
	return r.HandlerOther(r, "Fields", fields).([]*Field)
}

// FieldByName implementation of Scope interface.
func (r *Scope) FieldByName(name string) (field *Field, ok bool) {
	rf, rs := r.HandlerOthers(r, "FieldByName", field, ok, name)
	return rf.(*Field), rs.(bool)
}

// PrimaryFields implementation of Scope interface.
func (r *Scope) PrimaryFields() (fields []*Field) {
	return r.HandlerOther(r, "PrimaryFields", fields).([]*Field)
}

// PrimaryField implementation of Scope interface.
func (r *Scope) PrimaryField() (field *Field) {
	return r.HandlerOther(r, "PrimaryField", field).(*Field)
}

// PrimaryKey implementation of Scope interface.
func (r *Scope) PrimaryKey() (o string) {
	return r.HandlerOther(r, "PrimaryKey", o).(string)
}

// PrimaryKeyZero implementation of Scope interface.
func (r *Scope) PrimaryKeyZero() (o bool) {
	return r.HandlerOther(r, "PrimaryKeyZero", o).(bool)
}

// PrimaryKeyValue implementation of Scope interface.
func (r *Scope) PrimaryKeyValue() (o interface{}) {
	return r.HandlerOther(r, "PrimaryKeyValue", o)
}

// HasColumn implementation of Scope interface.
func (r *Scope) HasColumn(column string) (o bool) {
	return r.HandlerOther(r, "HasColumn", o, column).(bool)
}

// SetColumn implementation of Scope interface.
func (r *Scope) SetColumn(column interface{}, value interface{}) (o error) {
	return r.HandlerOther(r, "SetColumn", o, column, value).(error)
}

// CallMethod implementation of Scope interface.
func (r *Scope) CallMethod(methodName string) {
	r.Handler(r, "CallMethod", methodName)
}

// AddToVars implementation of Scope interface.
func (r *Scope) AddToVars(value interface{}) (o string) {
	return r.HandlerOther(r, "AddToVars", o, value).(string)
}

// SelectAttrs implementation of Scope interface.
func (r *Scope) SelectAttrs() (o []string) {
	return r.HandlerOther(r, "SelectAttrs", o).([]string)
}

// OmitAttrs implementation of Scope interface.
func (r *Scope) OmitAttrs() (o []string) {
	return r.HandlerOther(r, "OmitAttrs", o).([]string)
}

// TableName implementation of Scope interface.
func (r *Scope) TableName() (o string) {
	return r.HandlerOther(r, "TableName", o).(string)
}

// QuotedTableName implementation of Scope interface.
func (r *Scope) QuotedTableName() (o string) {
	return r.HandlerOther(r, "QuotedTableName", o).(string)
}

// CombinedConditionSql implementation of Scope interface.
func (r *Scope) CombinedConditionSql() (o string) {
	return r.HandlerOther(r, "CombinedConditionSql", o).(string)
}

// Raw implementation of Scope interface.
func (r *Scope) Raw(sql string) (o *Scope) {
	return r.HandlerOther(r, "Raw", o, sql).(*Scope)
}

// Exec implementation of Scope interface.
func (r *Scope) Exec() (o *Scope) {
	return r.HandlerOther(r, "Exec", o).(*Scope)
}

// Set implementation of Scope interface.
func (r *Scope) Set(name string, value interface{}) (o *Scope) {
	return r.HandlerOther(r, "Set", o, name, value).(*Scope)
}

// Get implementation of Scope interface.
func (r *Scope) Get(name string) (first interface{}, second bool) {
	rf, rs := r.HandlerOthers(r, "Get", first, second, name)
	return rf.(*Field), rs.(bool)
}

// InstanceID implementation of Scope interface.
func (r *Scope) InstanceID() (o string) {
	return r.HandlerOther(r, "InstanceID", o).(string)
}

// InstanceSet implementation of Scope interface.
func (r *Scope) InstanceSet(name string, value interface{}) (o *Scope) {
	return r.HandlerOther(r, "InstanceSet", o, name, value).(*Scope)
}

// InstanceGet implementation of Scope interface.
func (r *Scope) InstanceGet(name string) (first interface{}, second bool) {
	rf, rs := r.HandlerOthers(r, "InstanceGet", first, second, name)
	return rf.(*Field), rs.(bool)
}

// Begin implementation of Scope interface.
func (r *Scope) Begin() (o *Scope) {
	return r.HandlerOther(r, "Begin", o).(*Scope)
}

// CommitOrRollback implementation of Scope interface.
func (r *Scope) CommitOrRollback() (o *Scope) {
	return r.HandlerOther(r, "CommitOrRollback", o).(*Scope)
}
