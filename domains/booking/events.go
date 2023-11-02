package booking

import (
	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.booking/config"
)

type Events interface {
	Created(CreatedEvent)
	Validated(ValidatedEvent)
}

type (
	CreatedEvent struct {
		Entity *Entity `json:"entity"`
	}
	ValidatedEvent struct {
		Entity *Entity `json:"entity"`
	}
)

type bookingEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(cnf EventConfig) Events {
	return &bookingEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
	}
}

func (e bookingEvents) Created(event CreatedEvent) {}

func (e bookingEvents) Validated(event ValidatedEvent) {}
