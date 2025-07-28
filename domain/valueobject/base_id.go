package valueobject

// ID represents a base ID with common functionality
type WithID[T comparable] struct {
	ID T
}

// GetValue returns the ID value
func (id *WithID[T]) GetValue() T {
	return id.ID
}

// SetValue sets the ID value
func (id *WithID[T]) SetValue(value T) {
	id.ID = value
}

// Equals checks if two IDs are equal
func (id *WithID[T]) Equals(other *WithID[T]) bool {
	return id.ID == other.ID
}
