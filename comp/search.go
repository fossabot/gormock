// Copyright © 2018-2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comp

// Search a type that combines the capabilities of a mock and an interface
// that is almost fully compatible with gorm.
type Search struct {
	// Mock fields definition
	Foundation

	// Gorm fields definition
	Unscoped bool
}

// Where implementation of gorm interface.
func (r *Search) Where(query interface{}, values ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Where", o, query, values).(*Search)
}

// Not implementation of gorm interface.
func (r *Search) Not(query interface{}, values ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Not", o, query, values).(*Search)
}

// Or implementation of gorm interface.
func (r *Search) Or(query interface{}, values ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Or", o, query, values).(*Search)
}

// Attrs implementation of gorm interface.
func (r *Search) Attrs(attrs ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Attrs", o, attrs).(*Search)
}

// Assign implementation of gorm interface.
func (r *Search) Assign(attrs ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Assign", o, attrs).(*Search)
}

// Order implementation of gorm interface.
func (r *Search) Order(value interface{}, reorder ...bool) (o *Search) {
	return r.HandlerOther(r, "Order", o, value, reorder).(*Search)
}

// Select implementation of gorm interface.
func (r *Search) Select(query interface{}, args ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Select", o, query, args).(*Search)
}

// Omit implementation of gorm interface.
func (r *Search) Omit(columns ...string) (o *Search) {
	return r.HandlerOther(r, "Omit", o, columns).(*Search)
}

// Limit implementation of gorm interface.
func (r *Search) Limit(limit interface{}) (o *Search) {
	return r.HandlerOther(r, "Limit", o, limit).(*Search)
}

// Offset implementation of gorm interface.
func (r *Search) Offset(offset interface{}) (o *Search) {
	return r.HandlerOther(r, "Offset", o, offset).(*Search)
}

// Group implementation of gorm interface.
func (r *Search) Group(query string) (o *Search) {
	return r.HandlerOther(r, "Group", o, query).(*Search)
}

// Having implementation of gorm interface.
func (r *Search) Having(query interface{}, values ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Having", o, query, values).(*Search)
}

// Joins implementation of gorm interface.
func (r *Search) Joins(query string, values ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Joins", o, query, values).(*Search)
}

// Preload implementation of gorm interface.
func (r *Search) Preload(schema string, values ...interface{}) (o *Search) {
	return r.HandlerOther(r, "Preload", o, schema, values).(*Search)
}

// Raw implementation of gorm interface.
func (r *Search) Raw(b bool) (o *Search) {
	return r.HandlerOther(r, "Raw", o, b).(*Search)
}

// Table implementation of gorm interface.
func (r *Search) Table(name string) (o *Search) {
	return r.HandlerOther(r, "Table", o, name).(*Search)
}
