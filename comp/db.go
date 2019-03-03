// Copyright © 2018-2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comp

import (
	"database/sql"
)

// DB a type that combines the capabilities of a mock and an interface
// that is almost fully compatible with gorm.
type DB struct {
	// Mock fields definition
	Foundation

	// Gorm fields definition
	Value        interface{}
	Error        error
	RowsAffected int64
}

// clone the method for cloning an object, which makes a copy of the object
// and sets a pointer to the parent object, is designed to preserve some of
// the behavior of the original object and reduce problems with race condition.
func (r *DB) clone() (o *DB) {
	o = &DB{}
	*o = *r
	return o
}

// New implementation of gorm interface.
func (r *DB) New() (clone *DB) {
	return r.clone().HandlerOther(r, "New", nil).(*DB)
}

// Close implementation of gorm interface.
func (r *DB) Close() (err error) {
	return r.HandlerOther(r, "Close", err).(error)
}

// DB implementation of gorm interface.
func (r *DB) DB() (db *sql.DB) {
	return r.HandlerOther(r, "DB", db).(*sql.DB)
}

// CommonDB implementation of gorm interface.
func (r *DB) CommonDB() (sqlCommon SQLCommon) {
	return r.HandlerOther(r, "CommonDB", sqlCommon).(SQLCommon)
}

// Dialect implementation of gorm interface.
func (r *DB) Dialect() (dialect Dialect) {
	return r.HandlerOther(r, "Dialect", dialect).(Dialect)
}

// Callback implementation of gorm interface.
func (r *DB) Callback() (callback *Callback) {
	return r.HandlerOther(r, "Callback", callback).(*Callback)
}

// SetLogger implementation of gorm interface.
func (r *DB) SetLogger(log interface{ Print(v ...interface{}) }) {
	r.Handler(r, "SetLogger")
}

// LogMode implementation of gorm interface.
func (r *DB) LogMode(enable bool) (o *DB) {
	return r.clone().HandlerOther(r, "LogMode", o, enable).(*DB)
}

// BlockGlobalUpdate implementation of gorm interface.
func (r *DB) BlockGlobalUpdate(enable bool) (o *DB) {
	return r.clone().HandlerOther(r, "BlockGlobalUpdate", o, enable).(*DB)
}

// HasBlockGlobalUpdate implementation of gorm interface.
func (r *DB) HasBlockGlobalUpdate() (ok bool) {
	return r.HandlerOther(r, "HasBlockGlobalUpdate", ok).(bool)
}

// SingularTable implementation of gorm interface.
func (r *DB) SingularTable(enable bool) {
	r.Handler(r, "SingularTable", enable)
}

// NewScope implementation of gorm interface.
func (r *DB) NewScope(value interface{}) (scope *Scope) {
	return r.clone().HandlerOther(r, "NewScope", scope, value).(*Scope)
}

// QueryExpr implementation of gorm interface.
func (r *DB) QueryExpr() interface{} { // *expr
	return r.HandlerOther(r, "QueryExpr", nil)
}

// SubQuery implementation of gorm interface.
func (r *DB) SubQuery() interface{} { // *expr
	return r.HandlerOther(r, "SubQuery", nil)
}

// Where implementation of gorm interface.
func (r *DB) Where(query interface{}, args ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Where", o, query, args).(*DB)
}

// Or implementation of gorm interface.
func (r *DB) Or(query interface{}, args ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Or", o, query, args).(*DB)
}

// Not implementation of gorm interface.
func (r *DB) Not(query interface{}, args ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Not", o, query, args).(*DB)
}

// Limit implementation of gorm interface.
func (r *DB) Limit(limit interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Limit", o, limit).(*DB)
}

// Offset implementation of gorm interface.
func (r *DB) Offset(offset interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Offset", o, offset).(*DB)
}

// Order implementation of gorm interface.
func (r *DB) Order(value interface{}, reorder ...bool) (o *DB) {
	return r.clone().HandlerOther(r, "Order", o, value, reorder).(*DB)
}

// Select implementation of gorm interface.
func (r *DB) Select(query interface{}, args ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Select", o, query, args).(*DB)
}

// Omit implementation of gorm interface.
func (r *DB) Omit(columns ...string) (o *DB) {
	return r.clone().HandlerOther(r, "Omit", o, columns).(*DB)
}

// Group implementation of gorm interface.
func (r *DB) Group(query string) (o *DB) {
	return r.clone().HandlerOther(r, "Group", o, query).(*DB)
}

// Having implementation of gorm interface.
func (r *DB) Having(query interface{}, values ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Having", o, query, values).(*DB)
}

// Joins implementation of gorm interface.
func (r *DB) Joins(query string, args ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Joins", o, query, args).(*DB)
}

// Scopes implementation of gorm interface.
func (r *DB) Scopes(funcs ...func(*DB) *DB) (o *DB) {
	return r.clone().HandlerOther(r, "Scopes", o, funcs).(*DB)
}

// Unscoped implementation of gorm interface.
func (r *DB) Unscoped() (o *DB) {
	return r.clone().HandlerOther(r, "Unscoped", o).(*DB)
}

// Attrs implementation of gorm interface.
func (r *DB) Attrs(attrs ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Attrs", o, attrs).(*DB)
}

// Assign implementation of gorm interface.
func (r *DB) Assign(attrs ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Assign", o, attrs).(*DB)
}

// First implementation of gorm interface.
func (r *DB) First(out interface{}, where ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "First", o, out, where).(*DB)
}

// Take implementation of gorm interface.
func (r *DB) Take(out interface{}, where ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Take", o, out, where).(*DB)
}

// Last implementation of gorm interface.
func (r *DB) Last(out interface{}, where ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Last", o, out, where).(*DB)
}

// Find implementation of gorm interface.
func (r *DB) Find(out interface{}, where ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Find", o, out, where).(*DB)
}

// Scan implementation of gorm interface.
func (r *DB) Scan(dest interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Scan", o, dest).(*DB)
}

// Row implementation of gorm interface.
func (r *DB) Row() (row *sql.Row) {
	return r.HandlerOther(r, "Row", row).(*sql.Row)
}

// Rows implementation of gorm interface.
func (r *DB) Rows() (rows *sql.Rows, err error) {
	resultFirst, resultSecond := r.HandlerOthers(r, "Rows", rows, err)
	return resultFirst.(*sql.Rows), resultSecond.(error)
}

// ScanRows implementation of gorm interface.
func (r *DB) ScanRows(rows *sql.Rows, result interface{}) (err error) {
	return r.HandlerOther(r, "ScanRows", err).(error)
}

// Pluck implementation of gorm interface.
func (r *DB) Pluck(column string, value interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Pluck", o, column, value).(*DB)
}

// Count implementation of gorm interface.
func (r *DB) Count(value interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Count", o, value).(*DB)
}

// Related implementation of gorm interface.
func (r *DB) Related(value interface{}, foreignKeys ...string) (o *DB) {
	return r.clone().HandlerOther(r, "Related", o, value, foreignKeys).(*DB)
}

// FirstOrInit implementation of gorm interface.
func (r *DB) FirstOrInit(out interface{}, where ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "FirstOrInit", o, out, where).(*DB)
}

// FirstOrCreate implementation of gorm interface.
func (r *DB) FirstOrCreate(out interface{}, where ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "FirstOrCreate", o, out, where).(*DB)
}

// Update implementation of gorm interface.
func (r *DB) Update(attrs ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Update", o, attrs).(*DB)
}

// Updates implementation of gorm interface.
func (r *DB) Updates(values interface{}, ignore ...bool) (o *DB) {
	return r.clone().HandlerOther(r, "Updates", o, values, ignore).(*DB)
}

// UpdateColumn implementation of gorm interface.
func (r *DB) UpdateColumn(attrs ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "UpdateColumn", o, attrs).(*DB)
}

// UpdateColumns implementation of gorm interface.
func (r *DB) UpdateColumns(values interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "UpdateColumns", o, values).(*DB)
}

// Save implementation of gorm interface.
func (r *DB) Save(value interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Save", o, value).(*DB)
}

// Create implementation of gorm interface.
func (r *DB) Create(value interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Create", o).(*DB)
}

// Delete implementation of gorm interface.
func (r *DB) Delete(value interface{}, where ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Delete", o, value, where).(*DB)
}

// Raw implementation of gorm interface.
func (r *DB) Raw(sql string, values ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Raw", o, sql, values).(*DB)
}

// Exec implementation of gorm interface.
func (r *DB) Exec(sql string, values ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Exec", o, sql, values).(*DB)
}

// Model implementation of gorm interface.
func (r *DB) Model(value interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Model", o, value).(*DB)
}

// Table implementation of gorm interface.
func (r *DB) Table(name string) (o *DB) {
	return r.clone().HandlerOther(r, "Table", o, name).(*DB)
}

// Debug implementation of gorm interface.
func (r *DB) Debug() (o *DB) {
	return r.clone().HandlerOther(r, "Debug", o).(*DB)
}

// Begin implementation of gorm interface.
func (r *DB) Begin() (o *DB) {
	return r.clone().HandlerOther(r, "Begin", o).(*DB)
}

// Commit implementation of gorm interface.
func (r *DB) Commit() (o *DB) {
	return r.clone().HandlerOther(r, "Commit", o).(*DB)
}

// Rollback implementation of gorm interface.
func (r *DB) Rollback() (o *DB) {
	return r.clone().HandlerOther(r, "Rollback", o).(*DB)
}

// NewRecord implementation of gorm interface.
func (r *DB) NewRecord(value interface{}) (is bool) {
	return r.HandlerOther(r, "NewRecord", is, value).(bool)
}

// RecordNotFound implementation of gorm interface.
func (r *DB) RecordNotFound() (ok bool) {
	return r.HandlerOther(r, "RecordNotFound", ok).(bool)
}

// CreateTable implementation of gorm interface.
func (r *DB) CreateTable(models ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "CreateTable", o, models).(*DB)
}

// DropTable implementation of gorm interface.
func (r *DB) DropTable(values ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "DropTable", o, values).(*DB)
}

// DropTableIfExists implementation of gorm interface.
func (r *DB) DropTableIfExists(values ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "DropTableIfExists", o, values).(*DB)
}

// HasTable implementation of gorm interface.
func (r *DB) HasTable(value interface{}) (ok bool) {
	return r.HandlerOther(r, "HasTable", ok, value).(bool)
}

// AutoMigrate implementation of gorm interface.
func (r *DB) AutoMigrate(values ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "AutoMigrate", o, values).(*DB)
}

// ModifyColumn implementation of gorm interface.
func (r *DB) ModifyColumn(column string, typ string) (o *DB) {
	return r.clone().HandlerOther(r, "ModifyColumn", o, column, typ).(*DB)
}

// DropColumn implementation of gorm interface.
func (r *DB) DropColumn(column string) (o *DB) {
	return r.clone().HandlerOther(r, "DropColumn", o, column).(*DB)
}

// AddIndex implementation of gorm interface.
func (r *DB) AddIndex(indexName string, columns ...string) (o *DB) {
	return r.clone().HandlerOther(r, "AddIndex", o, indexName, columns).(*DB)
}

// AddUniqueIndex implementation of gorm interface.
func (r *DB) AddUniqueIndex(name string, columns ...string) (o *DB) {
	return r.clone().HandlerOther(r, "AddUniqueIndex", o, name, columns).(*DB)
}

// RemoveIndex implementation of gorm interface.
func (r *DB) RemoveIndex(indexName string) (o *DB) {
	return r.clone().HandlerOther(r, "RemoveIndex", o, indexName).(*DB)
}

// AddForeignKey implementation of gorm interface.
func (r *DB) AddForeignKey(field, dest, onDelete, onUpdate string) (o *DB) {
	return r.clone().HandlerOther(
		r, "AddForeignKey", o,
		field, dest, onDelete, onUpdate,
	).(*DB)
}

// RemoveForeignKey implementation of gorm interface.
func (r *DB) RemoveForeignKey(field string, dest string) (o *DB) {
	return r.clone().HandlerOther(r, "RemoveForeignKey", o, field, dest).(*DB)
}

// Association implementation of gorm interface.
func (r *DB) Association(column string) (o *Association) {
	return r.clone().HandlerOther(r, "Association", o, column).(*Association)
}

// Preload implementation of gorm interface.
func (r *DB) Preload(column string, conditions ...interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Preload", o, column, conditions).(*DB)
}

// Set implementation of gorm interface.
func (r *DB) Set(name string, value interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "Set", o, name, value).(*DB)
}

// InstantSet implementation of gorm interface.
func (r *DB) InstantSet(name string, value interface{}) (o *DB) {
	return r.clone().HandlerOther(r, "InstantSet", o, name, value).(*DB)
}

// Get implementation of gorm interface.
func (r *DB) Get(name string) (value interface{}, ok bool) {
	resultFirst, resultSecond := r.HandlerOthers(r, "Get", value, ok, name)
	return resultFirst, resultSecond.(bool)
}

// SetJoinTableHandler implementation of gorm interface.
func (r *DB) SetJoinTableHandler(
	source interface{},
	column string,
	handler JoinTableHandlerInterface,
) {
	r.Handler(r, "SetJoinTableHandler", source, column, handler)
}

// AddError implementation of gorm interface.
func (r *DB) AddError(err error) (e error) {
	return r.HandlerOther(r, "AddError", e, err).(error)
}

// GetErrors implementation of gorm interface.
func (r *DB) GetErrors() (errs []error) {
	return r.HandlerOther(r, "GetErrors", errs).([]error)
}
