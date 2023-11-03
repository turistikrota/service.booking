package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.booking/app"
	"github.com/turistikrota/service.booking/app/command"
	"github.com/turistikrota/service.booking/config"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Validator   *validation.Validator
	MongoDB     *mongo.DB
	CacheSrv    cache.Service
}

func NewApplication(cnf Config) app.Application {
	bookingFactory := booking.NewFactory()
	bookingRepo := booking.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Booking.Collection), bookingFactory)
	bookingEvents := booking.NewEvents(booking.EventConfig{
		Topics:    cnf.App.Topics,
		Publisher: cnf.EventEngine,
	})

	return app.Application{
		Commands: app.Commands{
			BookingCreate: command.NewBookingCreateHandler(bookingFactory, bookingRepo, bookingEvents),
		},
		Queries: app.Queries{},
	}
}
