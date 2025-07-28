package entity

// AggregateRoot represents an aggregate root with domain events
type AggregateRoot[T comparable] struct {
	Entity[T]
}
