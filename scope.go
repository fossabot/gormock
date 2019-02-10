package gormock

import (
	"reflect"

	"github.com/stretchr/testify/mock"
)

// Scope a type that combines the capabilities of a mock and an interface that
// is almost fully compatible with gorm.Scope.
type Scope struct {
	// Mock fields definition
	Foundation
	*mock.Mock
	db *DB

	// Scope public fields
	Search  Search
	Value   interface{}
	SQL     string
	SQLVars []interface{}
}

// IndirectValue implementation of Scope interface.
func (s *Scope) IndirectValue() (value reflect.Value) {
	return s.HandlerOther(s, "IndirectValue", value).(reflect.Value)
}

// New implementation of Scope interface.
func (s *Scope) New(value interface{}) (outcome *Scope) {
	return s.HandlerOther(s, "New", outcome, value).(*Scope)
}

// DB implementation of Scope interface.
func (s *Scope) DB() *DB {
	return s.db
}

// NewDB implementation of Scope interface.
func (s *Scope) NewDB() (db *DB) {
	return s.HandlerOther(s, "NewDB", db).(*DB)
}

// SQLDB implementation of Scope interface.
func (s *Scope) SQLDB() (sqlCommon SQLCommon) {
	return s.HandlerOther(s, "SQLDB", sqlCommon).(SQLCommon)
}

// Dialect implementation of Scope interface.
func (s *Scope) Dialect() (dialect Dialect) {
	return s.HandlerOther(s, "Dialect", dialect).(Dialect)
}

// Quote implementation of Scope interface.
func (s *Scope) Quote(str string) (outcome string) {
	return s.HandlerOther(s, "Quote", outcome).(string)
}

// Err implementation of Scope interface.
func (s *Scope) Err(err error) (outcome error) {
	return s.HandlerOther(s, "Err", outcome, err).(error)
}

// HasError implementation of Scope interface.
func (s *Scope) HasError() (outcome bool) {
	return s.HandlerOther(s, "HasError", outcome).(bool)
}

// Log implementation of Scope interface.
func (s *Scope) Log(v ...interface{}) {
	s.Handler(s, "Log", v)
}

// SkipLeft implementation of Scope interface.
func (s *Scope) SkipLeft() {
	s.Handler(s, "SkipLeft")
}

// Fields implementation of Scope interface.
func (s *Scope) Fields() (fields []*Field) {
	return s.HandlerOther(s, "Fields", fields).([]*Field)
}

// FieldByName implementation of Scope interface.
func (s *Scope) FieldByName(name string) (field *Field, ok bool) {
	rf, rs := s.HandlerOthers(s, "FieldByName", field, ok, name)
	return rf.(*Field), rs.(bool)
}

// PrimaryFields implementation of Scope interface.
func (s *Scope) PrimaryFields() (fields []*Field) {
	return s.HandlerOther(s, "PrimaryFields", fields).([]*Field)
}

// PrimaryField implementation of Scope interface.
func (s *Scope) PrimaryField() (field *Field) {
	return s.HandlerOther(s, "PrimaryField", field).(*Field)
}

// PrimaryKey implementation of Scope interface.
func (s *Scope) PrimaryKey() (outcome string) {
	return s.HandlerOther(s, "PrimaryKey", outcome).(string)
}

// PrimaryKeyZero implementation of Scope interface.
func (s *Scope) PrimaryKeyZero() (outcome bool) {
	return s.HandlerOther(s, "PrimaryKeyZero", outcome).(bool)
}

// PrimaryKeyValue implementation of Scope interface.
func (s *Scope) PrimaryKeyValue() (outcome interface{}) {
	return s.HandlerOther(s, "PrimaryKeyValue", outcome)
}

// HasColumn implementation of Scope interface.
func (s *Scope) HasColumn(column string) (outcome bool) {
	return s.HandlerOther(s, "HasColumn", outcome, column).(bool)
}

// SetColumn implementation of Scope interface.
func (s *Scope) SetColumn(column interface{}, value interface{}) (outcome error) {
	return s.HandlerOther(s, "SetColumn", outcome, column, value).(error)
}

// CallMethod implementation of Scope interface.
func (s *Scope) CallMethod(methodName string) {
	s.Handler(s, "CallMethod", methodName)
}

// AddToVars implementation of Scope interface.
func (s *Scope) AddToVars(value interface{}) (outcome string) {
	return s.HandlerOther(s, "AddToVars", outcome, value).(string)
}

// SelectAttrs implementation of Scope interface.
func (s *Scope) SelectAttrs() (outcome []string) {
	return s.HandlerOther(s, "SelectAttrs", outcome).([]string)
}

// OmitAttrs implementation of Scope interface.
func (s *Scope) OmitAttrs() (outcome []string) {
	return s.HandlerOther(s, "OmitAttrs", outcome).([]string)
}

// TableName implementation of Scope interface.
func (s *Scope) TableName() (outcome string) {
	return s.HandlerOther(s, "TableName", outcome).(string)
}

// QuotedTableName implementation of Scope interface.
func (s *Scope) QuotedTableName() (outcome string) {
	return s.HandlerOther(s, "QuotedTableName", outcome).(string)
}

// CombinedConditionSql implementation of Scope interface.
func (s *Scope) CombinedConditionSql() (outcome string) {
	return s.HandlerOther(s, "CombinedConditionSql", outcome).(string)
}

// Raw implementation of Scope interface.
func (s *Scope) Raw(sql string) (outcome *Scope) {
	return s.HandlerOther(s, "Raw", outcome, sql).(*Scope)
}

// Exec implementation of Scope interface.
func (s *Scope) Exec() (outcome *Scope) {
	return s.HandlerOther(s, "Exec", outcome).(*Scope)
}

// Set implementation of Scope interface.
func (s *Scope) Set(name string, value interface{}) (outcome *Scope) {
	return s.HandlerOther(s, "Set", outcome, name, value).(*Scope)
}

// Get implementation of Scope interface.
func (s *Scope) Get(name string) (first interface{}, second bool) {
	rf, rs := s.HandlerOthers(s, "Get", first, second, name)
	return rf.(*Field), rs.(bool)
}

// InstanceID implementation of Scope interface.
func (s *Scope) InstanceID() (outcome string) {
	return s.HandlerOther(s, "InstanceID", outcome).(string)
}

// InstanceSet implementation of Scope interface.
func (s *Scope) InstanceSet(name string, value interface{}) (outcome *Scope) {
	return s.HandlerOther(s, "InstanceSet", outcome, name, value).(*Scope)
}

// InstanceGet implementation of Scope interface.
func (s *Scope) InstanceGet(name string) (first interface{}, second bool) {
	rf, rs := s.HandlerOthers(s, "InstanceGet", first, second, name)
	return rf.(*Field), rs.(bool)
}

// Begin implementation of Scope interface.
func (s *Scope) Begin() (outcome *Scope) {
	return s.HandlerOther(s, "Begin", outcome).(*Scope)
}

// CommitOrRollback implementation of Scope interface.
func (s *Scope) CommitOrRollback() (outcome *Scope) {
	return s.HandlerOther(s, "CommitOrRollback", outcome).(*Scope)
}
