package event_stream

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/server"
	"github.com/turistikrota/service.booking/app"
	"github.com/turistikrota/service.booking/config"
)

type srv struct {
	app    app.Application
	topics config.Topics
	engine events.Engine
}

type Config struct {
	App    app.Application
	Engine events.Engine
	Topics config.Topics
}

func New(config Config) server.Server {
	return srv{
		app:    config.App,
		engine: config.Engine,
		topics: config.Topics,
	}
}

func (s srv) Listen() error {
	err := s.engine.Subscribe(s.topics.Booking.ValidationSuccess, s.OnBookingValidationSucceed)
	if err != nil {
		return err
	}
	err = s.engine.Subscribe(s.topics.Booking.ValidationFail, s.OnBookingValidationFail)
	if err != nil {
		return err
	}
	err = s.engine.Subscribe(s.topics.Booking.PaySuccess, s.OnBookingPaySuccess)
	if err != nil {
		return err
	}
	err = s.engine.Subscribe(s.topics.Booking.PayTimeout, s.OnBookingPayTimeout)
	return err
}
