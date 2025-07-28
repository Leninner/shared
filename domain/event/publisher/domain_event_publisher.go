package publisher

import "github.com/leninner/shared/domain/event"

type DomainEventPublisher[T event.DomainEvent] interface {
	Publish(domainEvent T) error
}
