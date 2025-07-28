package entity

import "github.com/leninner/shared/domain/valueobject"

// Entity represents a base entity with common functionality
type Entity[T comparable] struct {
	valueobject.WithID[T]
}

// GetID returns the entity's ID
func (e *Entity[T]) GetID() T {
	return e.WithID.GetValue()
}

// SetID sets the entity's ID
func (e *Entity[T]) SetID(id T) {
	e.WithID.SetValue(id)
}