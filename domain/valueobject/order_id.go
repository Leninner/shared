package valueobject

import "github.com/google/uuid"

type OrderID struct {
	WithID[uuid.UUID]
}

func NewOrderID() OrderID {
	return OrderID{
		WithID: WithID[uuid.UUID]{
			ID: uuid.New(),
		},
	}
}

func NewOrderIDFromUUID(id uuid.UUID) OrderID {
	return OrderID{
		WithID: WithID[uuid.UUID]{
			ID: id,
		},
	}
}
