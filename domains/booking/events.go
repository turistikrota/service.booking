package booking

import (
	"time"

	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.booking/config"
)

type Events interface {
	Created(CreatedEvent)
	PayPending(PayPendingEvent)
}

type (
	CreatedEvent struct {
		BookingUUID string    `json:"booking_uuid"`
		PostUUID    string    `json:"post_uuid"`
		People      *People   `json:"people"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
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
	_ = e.publisher.Publish(e.topics.Booking.PayPending, event)
}
