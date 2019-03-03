// Copyright © 2018-2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comp

// CallbackProcessor a type that combines the capabilities of a mock
// and an interface that is almost fully compatible with gorm.
type CallbackProcessor struct {
	// Mock fields definition
	Foundation
}

// After implementation of gorm interface.
func (r *CallbackProcessor) After(callbackName string) (o *CallbackProcessor) {
	return r.HandlerOther(r, "After", o, callbackName).(*CallbackProcessor)
}

// Before implementation of gorm interface.
func (r *CallbackProcessor) Before(callbackName string) (o *CallbackProcessor) {
	return r.HandlerOther(r, "Before", o, callbackName).(*CallbackProcessor)
}

// Register implementation of gorm interface.
func (r *CallbackProcessor) Register(name string, callback func(*Scope)) {
	r.Handler(r, "Register", name, callback)
}

// Remove implementation of gorm interface.
func (r *CallbackProcessor) Remove(callbackName string) {
	r.Handler(r, "Remove", callbackName)
}

// Replace implementation of gorm interface.
func (r *CallbackProcessor) Replace(name string, callback func(*Scope)) {
	r.Handler(r, "Replace", name, callback)
}

// Get implementation of gorm interface.
func (r *CallbackProcessor) Get(callbackName string) (o func(*Scope)) {
	return r.HandlerOther(r, "Get", o, callbackName).(func(*Scope))
}
