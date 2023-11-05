package booking

import (
	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.booking/config"
)

type Events interface {
	Created(CreatedEvent)
	PayPending(PayPendingEvent)
}

type (
	CreatedEvent struct {
		Entity *Entity `json:"entity"`
	}
	PayPendingEvent struct {
		BookingUUID string `json:"booking_uuid"`
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

func (e bookingEvents) Created(event CreatedEvent) {
	_ = e.publisher.Publish(e.topics.Booking.ValidationStart, event)
}

func (e bookingEvents) PayPending(event PayPendingEvent) {
	_ = e.publisher.Publish(e.topics.Booking.ValidationStart, event)
}
