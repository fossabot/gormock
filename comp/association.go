package comp

// Association a type that combines the capabilities of a mock
// and an interface that is almost fully compatible with gorm.
type Association struct {
	// Mock fields definition.
	Foundation

	// Gorm fields definition.
	Error error
}

// Find implementation of gorm interface.
func (r Association) Find(value interface{}) (o *Association) {
	return r.HandlerOther(r, "Find", o, value).(*Association)
}

// Append implementation of gorm interface.
func (r Association) Append(values ...interface{}) (o *Association) {
	return r.HandlerOther(r, "Append", o, values).(*Association)
}

// Replace implementation of gorm interface.
func (r Association) Replace(values ...interface{}) (o *Association) {
	return r.HandlerOther(r, "Replace", o, values).(*Association)
}

// Delete implementation of gorm interface.
func (r Association) Delete(values ...interface{}) (o *Association) {
	return r.HandlerOther(r, "Delete", o, values).(*Association)
}

// Clear implementation of gorm interface.
func (r Association) Clear() (o *Association) {
	return r.HandlerOther(r, "Clear", o).(*Association)
}

// Count implementation of gorm interface.
func (r Association) Count() (o int) {
	return r.HandlerOther(r, "Count", o).(int)
}
