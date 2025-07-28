package valueobject

import "github.com/google/uuid"

type ProductID struct {
	WithID[uuid.UUID]
}

func NewProductID() ProductID {
	return ProductID{
		WithID: WithID[uuid.UUID]{
			ID: uuid.New(),
		},
	}
}

func NewProductIDFromUUID(id *uuid.UUID) ProductID {
	return ProductID{WithID: WithID[uuid.UUID]{ID: *id}}
}