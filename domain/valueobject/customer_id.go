package valueobject

import "github.com/google/uuid"

type CustomerID struct {
	WithID[uuid.UUID]
}

func NewCustomerID() CustomerID {
	return CustomerID{
		WithID: WithID[uuid.UUID]{
			ID: uuid.New(),
		},
	}
}

func NewCustomerIDFromUUID(id *uuid.UUID) CustomerID {
	return CustomerID{WithID: WithID[uuid.UUID]{ID: *id}}
}