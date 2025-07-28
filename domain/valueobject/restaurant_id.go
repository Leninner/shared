package valueobject

import "github.com/google/uuid"

type RestaurantID struct {
	WithID[uuid.UUID]
}

func NewRestaurantID() RestaurantID {
	return RestaurantID{
		WithID: WithID[uuid.UUID]{
			ID: uuid.New(),
		},
	}
}

func NewRestaurantIDFromUUID(id *uuid.UUID) RestaurantID {
	return RestaurantID{WithID: WithID[uuid.UUID]{ID: *id}}
}