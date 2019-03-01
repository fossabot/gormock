package comp

// Callback a type that combines the capabilities of a mock
// and an interface that is almost fully compatible with gorm.
type Callback struct {
	// Mock fields definition
	Foundation
}

// Create implementation of gorm interface.
func (r *Callback) Create() (o *CallbackProcessor) {
	return r.HandlerOther(r, "Create", o).(*CallbackProcessor)
}

// Update implementation of gorm interface.
func (r *Callback) Update() (o *CallbackProcessor) {
	return r.HandlerOther(r, "Update", o).(*CallbackProcessor)
}

// Delete implementation of gorm interface.
func (r *Callback) Delete() (o *CallbackProcessor) {
	return r.HandlerOther(r, "Delete", o).(*CallbackProcessor)
}

// Query implementation of gorm interface.
func (r *Callback) Query() (o *CallbackProcessor) {
	return r.HandlerOther(r, "Query", o).(*CallbackProcessor)
}

// RowQuery implementation of gorm interface.
func (r *Callback) RowQuery() (o *CallbackProcessor) {
	return r.HandlerOther(r, "RowQuery", o).(*CallbackProcessor)
}
