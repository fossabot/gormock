package gormock

import (
	"database/sql"

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

// New implementation of gorm interface.
func (s *DB) New() (clone *DB) {
	return s.clone().HandlerOther(s, "New", nil).(*DB)
}

// Close implementation of gorm interface.
func (s *DB) Close() (err error) {
	return s.HandlerOther(s, "Close", err).(error)
}

// DB implementation of gorm interface.
func (s *DB) DB() (db *sql.DB) {
	return s.HandlerOther(s, "DB", db).(*sql.DB)
}

// CommonDB implementation of gorm interface.
func (s *DB) CommonDB() (sqlCommon SQLCommon) {
	return s.HandlerOther(s, "CommonDB", sqlCommon).(SQLCommon)
}

// Dialect implementation of gorm interface.
func (s *DB) Dialect() (dialect Dialect) {
	return s.HandlerOther(s, "Dialect", dialect).(Dialect)
}

// Callback implementation of gorm interface.
func (s *DB) Callback() (callback *Callback) {
	return s.HandlerOther(s, "Callback", callback).(*Callback)
}

// SetLogger implementation of gorm interface.
func (s *DB) SetLogger(log interface{ Print(v ...interface{}) }) {
	s.Handler(s, "SetLogger")
}

// LogMode implementation of gorm interface.
func (s *DB) LogMode(enable bool) (outcome *DB) {
	return s.clone().HandlerOther(s, "LogMode", outcome, enable).(*DB)
}

// BlockGlobalUpdate implementation of gorm interface.
func (s *DB) BlockGlobalUpdate(enable bool) (outcome *DB) {
	return s.clone().HandlerOther(s, "BlockGlobalUpdate", outcome, enable).(*DB)
}

// HasBlockGlobalUpdate implementation of gorm interface.
func (s *DB) HasBlockGlobalUpdate() (ok bool) {
	return s.HandlerOther(s, "HasBlockGlobalUpdate", ok).(bool)
}

// SingularTable implementation of gorm interface.
func (s *DB) SingularTable(enable bool) {
	s.Handler(s, "SingularTable")
}

// NewScope implementation of gorm interface.
func (s *DB) NewScope(value interface{}) (scope *Scope) {
	return s.clone().HandlerOther(s, "NewScope", scope, value).(*Scope)
}

// QueryExpr implementation of gorm interface.
func (s *DB) QueryExpr() interface{} { // *expr
	return s.HandlerOther(s, "QueryExpr", nil)
}

// SubQuery implementation of gorm interface.
func (s *DB) SubQuery() interface{} { // *expr
	return s.HandlerOther(s, "SubQuery", nil)
}

// Where implementation of gorm interface.
func (s *DB) Where(query interface{}, args ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Where", outcome, query, args).(*DB)
}

// Or implementation of gorm interface.
func (s *DB) Or(query interface{}, args ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Or", outcome, query, args).(*DB)
}

// Not implementation of gorm interface.
func (s *DB) Not(query interface{}, args ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Not", outcome, query, args).(*DB)
}

// Limit implementation of gorm interface.
func (s *DB) Limit(limit interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Limit", outcome, limit).(*DB)
}

// Offset implementation of gorm interface.
func (s *DB) Offset(offset interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Offset", outcome, offset).(*DB)
}

// Order implementation of gorm interface.
func (s *DB) Order(value interface{}, reorder ...bool) (outcome *DB) {
	return s.clone().HandlerOther(s, "Order", outcome, value, reorder).(*DB)
}

// Select implementation of gorm interface.
func (s *DB) Select(query interface{}, args ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Select", outcome, query, args).(*DB)
}

// Omit implementation of gorm interface.
func (s *DB) Omit(columns ...string) (outcome *DB) {
	return s.clone().HandlerOther(s, "Omit", outcome, columns).(*DB)
}

// Group implementation of gorm interface.
func (s *DB) Group(query string) (outcome *DB) {
	return s.clone().HandlerOther(s, "Group", outcome, query).(*DB)
}

// Having implementation of gorm interface.
func (s *DB) Having(query interface{}, values ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Having", outcome, query, values).(*DB)
}

// Joins implementation of gorm interface.
func (s *DB) Joins(query string, args ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Joins", outcome, query, args).(*DB)
}

// Scopes implementation of gorm interface.
func (s *DB) Scopes(funcs ...func(*DB) *DB) (outcome *DB) {
	return s.clone().HandlerOther(s, "Scopes", outcome, funcs).(*DB)
}

// Unscoped implementation of gorm interface.
func (s *DB) Unscoped() (outcome *DB) {
	return s.clone().HandlerOther(s, "Unscoped", outcome).(*DB)
}

// Attrs implementation of gorm interface.
func (s *DB) Attrs(attrs ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Attrs", outcome, attrs).(*DB)
}

// Assign implementation of gorm interface.
func (s *DB) Assign(attrs ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Assign", outcome, attrs).(*DB)
}

// First implementation of gorm interface.
func (s *DB) First(out interface{}, where ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "First", outcome, out, where).(*DB)
}

// Take implementation of gorm interface.
func (s *DB) Take(out interface{}, where ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Take", outcome, out, where).(*DB)
}

// Last implementation of gorm interface.
func (s *DB) Last(out interface{}, where ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Last", outcome, out, where).(*DB)
}

// Find implementation of gorm interface.
func (s *DB) Find(out interface{}, where ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Find", outcome, out, where).(*DB)
}

// Scan implementation of gorm interface.
func (s *DB) Scan(dest interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Scan", outcome, dest).(*DB)
}

// Row implementation of gorm interface.
func (s *DB) Row() (row *sql.Row) {
	return s.HandlerOther(s, "Row", row).(*sql.Row)
}

// Rows implementation of gorm interface.
func (s *DB) Rows() (rows *sql.Rows, err error) {
	resultFirst, resultSecond := s.HandlerOthers(s, "Rows", rows, err)
	return resultFirst.(*sql.Rows), resultSecond.(error)
}

// ScanRows implementation of gorm interface.
func (s *DB) ScanRows(rows *sql.Rows, result interface{}) (err error) {
	return s.HandlerOther(s, "ScanRows", err).(error)
}

// Pluck implementation of gorm interface.
func (s *DB) Pluck(column string, value interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Pluck", outcome, column, value).(*DB)
}

// Count implementation of gorm interface.
func (s *DB) Count(value interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Count", outcome, value).(*DB)
}

// Related implementation of gorm interface.
func (s *DB) Related(value interface{}, foreignKeys ...string) (outcome *DB) {
	return s.clone().HandlerOther(s, "Related", outcome, value, foreignKeys).(*DB)
}

// FirstOrInit implementation of gorm interface.
func (s *DB) FirstOrInit(out interface{}, where ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "FirstOrInit", outcome, out, where).(*DB)
}

// FirstOrCreate implementation of gorm interface.
func (s *DB) FirstOrCreate(out interface{}, where ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "FirstOrCreate", outcome, out, where).(*DB)
}

// Update implementation of gorm interface.
func (s *DB) Update(attrs ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Update", outcome, attrs).(*DB)
}

// Updates implementation of gorm interface.
func (s *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) (outcome *DB) {
	return s.clone().HandlerOther(s, "Updates", outcome, values, ignoreProtectedAttrs).(*DB)
}

// UpdateColumn implementation of gorm interface.
func (s *DB) UpdateColumn(attrs ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "UpdateColumn", outcome, attrs).(*DB)
}

// UpdateColumns implementation of gorm interface.
func (s *DB) UpdateColumns(values interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "UpdateColumns", outcome, values).(*DB)
}

// Save implementation of gorm interface.
func (s *DB) Save(value interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Save", outcome, value).(*DB)
}

// Create implementation of gorm interface.
func (s *DB) Create(value interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Create", outcome).(*DB)
}

// Delete implementation of gorm interface.
func (s *DB) Delete(value interface{}, where ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Delete", outcome, value, where).(*DB)
}

// Raw implementation of gorm interface.
func (s *DB) Raw(sql string, values ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Raw", outcome, sql, values).(*DB)
}

// Exec implementation of gorm interface.
func (s *DB) Exec(sql string, values ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Exec", outcome, sql, values).(*DB)
}

// Model implementation of gorm interface.
func (s *DB) Model(value interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Model", outcome, value).(*DB)
}

// Table implementation of gorm interface.
func (s *DB) Table(name string) (outcome *DB) {
	return s.clone().HandlerOther(s, "Table", outcome, name).(*DB)
}

// Debug implementation of gorm interface.
func (s *DB) Debug() (outcome *DB) {
	return s.clone().HandlerOther(s, "Debug", outcome).(*DB)
}

// Begin implementation of gorm interface.
func (s *DB) Begin() (outcome *DB) {
	return s.clone().HandlerOther(s, "Begin", outcome).(*DB)
}

// Commit implementation of gorm interface.
func (s *DB) Commit() (outcome *DB) {
	return s.clone().HandlerOther(s, "Commit", outcome).(*DB)
}

// Rollback implementation of gorm interface.
func (s *DB) Rollback() (outcome *DB) {
	return s.clone().HandlerOther(s, "Rollback", outcome).(*DB)
}

// NewRecord implementation of gorm interface.
func (s *DB) NewRecord(value interface{}) (is bool) {
	return s.HandlerOther(s, "NewRecord", is, value).(bool)
}

// RecordNotFound implementation of gorm interface.
func (s *DB) RecordNotFound() (ok bool) {
	return s.HandlerOther(s, "RecordNotFound", ok).(bool)
}

// CreateTable implementation of gorm interface.
func (s *DB) CreateTable(models ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "CreateTable", outcome, models).(*DB)
}

// DropTable implementation of gorm interface.
func (s *DB) DropTable(values ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "DropTable", outcome, values).(*DB)
}

// DropTableIfExists implementation of gorm interface.
func (s *DB) DropTableIfExists(values ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "DropTableIfExists", outcome, values).(*DB)
}

// HasTable implementation of gorm interface.
func (s *DB) HasTable(value interface{}) (ok bool) {
	return s.HandlerOther(s, "HasTable", ok, value).(bool)
}

// AutoMigrate implementation of gorm interface.
func (s *DB) AutoMigrate(values ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "AutoMigrate", outcome, values).(*DB)
}

// ModifyColumn implementation of gorm interface.
func (s *DB) ModifyColumn(column string, typ string) (outcome *DB) {
	return s.clone().HandlerOther(s, "ModifyColumn", outcome, column, typ).(*DB)
}

// DropColumn implementation of gorm interface.
func (s *DB) DropColumn(column string) (outcome *DB) {
	return s.clone().HandlerOther(s, "DropColumn", outcome, column).(*DB)
}

// AddIndex implementation of gorm interface.
func (s *DB) AddIndex(indexName string, columns ...string) (outcome *DB) {
	return s.clone().HandlerOther(s, "AddIndex", outcome, indexName, columns).(*DB)
}

// AddUniqueIndex implementation of gorm interface.
func (s *DB) AddUniqueIndex(indexName string, columns ...string) (outcome *DB) {
	return s.clone().HandlerOther(s, "AddUniqueIndex", outcome, indexName, columns).(*DB)
}

// RemoveIndex implementation of gorm interface.
func (s *DB) RemoveIndex(indexName string) (outcome *DB) {
	return s.clone().HandlerOther(s, "RemoveIndex", outcome, indexName).(*DB)
}

// AddForeignKey implementation of gorm interface.
func (s *DB) AddForeignKey(field, dest, onDelete, onUpdate string) (outcome *DB) {
	return s.clone().HandlerOther(s, "AddForeignKey", outcome, field, dest, onDelete, onUpdate).(*DB)
}

// RemoveForeignKey implementation of gorm interface.
func (s *DB) RemoveForeignKey(field string, dest string) (outcome *DB) {
	return s.clone().HandlerOther(s, "RemoveForeignKey", outcome, field, dest).(*DB)
}

// Association implementation of gorm interface.
func (s *DB) Association(column string) (outcome *Association) {
	return s.clone().HandlerOther(s, "Association", outcome, column).(*Association)
}

// Preload implementation of gorm interface.
func (s *DB) Preload(column string, conditions ...interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Preload", outcome, column, conditions).(*DB)
}

// Set implementation of gorm interface.
func (s *DB) Set(name string, value interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "Set", outcome, name, value).(*DB)
}

// InstantSet implementation of gorm interface.
func (s *DB) InstantSet(name string, value interface{}) (outcome *DB) {
	return s.clone().HandlerOther(s, "InstantSet", outcome, name, value).(*DB)
}

// Get implementation of gorm interface.
func (s *DB) Get(name string) (value interface{}, ok bool) {
	resultFirst, resultSecond := s.HandlerOthers(s, "Get", value, ok, name)
	return resultFirst, resultSecond.(bool)
}

// SetJoinTableHandler implementation of gorm interface.
func (s *DB) SetJoinTableHandler(
	source interface{},
	column string,
	handler JoinTableHandlerInterface,
) {
	s.Handler(s, "SetJoinTableHandler", source, column, handler)
}

// AddError implementation of gorm interface.
func (s *DB) AddError(err error) (e error) {
	return s.HandlerOther(s, "AddError", e, err).(error)
}

// GetErrors implementation of gorm interface.
func (s *DB) GetErrors() (errs []error) {
	return s.HandlerOther(s, "GetErrors", errs).([]error)
}
