package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.booking/app"
	"github.com/turistikrota/service.booking/app/command"
	"github.com/turistikrota/service.booking/app/query"
	"github.com/turistikrota/service.booking/config"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/domains/invite"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Validator   *validation.Validator
	MongoDB     *mongo.DB
	CacheSrv    cache.Service
	I18n        *i18np.I18n
}

func NewApplication(cnf Config) app.Application {
	bookingFactory := booking.NewFactory()
	bookingRepo := booking.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Booking.Collection), bookingFactory)
	bookingEvents := booking.NewEvents(booking.EventConfig{
		Topics:    cnf.App.Topics,
		Publisher: cnf.EventEngine,
	})

	inviteFactory := invite.NewFactory()
	inviteRepo := invite.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Invite.Collection), inviteFactory)
	inviteEvents := invite.NewEvents(invite.EventConfig{
		Topics:    cnf.App.Topics,
		Publisher: cnf.EventEngine,
		I18n:      cnf.I18n,
	})

	return app.Application{
		Commands: app.Commands{
			BookingCreate:            command.NewBookingCreateHandler(bookingFactory, bookingRepo, bookingEvents),
			BookingCancel:            command.NewBookingCancelHandler(bookingFactory, bookingRepo, bookingEvents),
			BookingMarkPrivate:       command.NewBookingMarkPrivateHandler(bookingFactory, bookingRepo),
			BookingPayTimeout:        command.NewBookingPayTimeoutHandler(bookingRepo),
			BookingPaySuccess:        command.NewBookingPaySuccessHandler(bookingRepo),
			BookingMarkPublic:        command.NewBookingMarkPublicHandler(bookingFactory, bookingRepo),
			BookingRemoveGuest:       command.NewBookingRemoveGuestHandler(bookingFactory, bookingRepo, bookingEvents),
			BookingGuestMarkPublic:   command.NewBookingGuestMarkPublicHandler(bookingFactory, bookingRepo, bookingEvents),
			BookingGuestMarkPrivate:  command.NewBookingGuestMarkPrivateHandler(bookingFactory, bookingRepo, bookingEvents),
			BookingValidationSucceed: command.NewBookingValidationSucceedHandler(bookingRepo, bookingEvents),
			BookingValidationFailed:  command.NewBookingValidationFailedHandler(bookingRepo),

			InviteCreate: command.NewInviteCreateHandler(bookingFactory, bookingRepo, inviteFactory, inviteRepo, inviteEvents),
			InviteUse:    command.NewInviteUseHandler(inviteFactory, inviteRepo, bookingRepo),
			InviteDelete: command.NewInviteDeleteHandler(inviteRepo),
		},
		Queries: app.Queries{
			BookingAdminList:         query.NewBookingAdminListHandler(bookingRepo),
			BookingAdminView:         query.NewBookingAdminViewHandler(bookingRepo),
			BookingCheckAvailability: query.NewBookingCheckAvailabilityHandler(bookingRepo),
			BookingListByBusiness:    query.NewBookingListByBusinessHandler(bookingRepo, cnf.CacheSrv),
			BookingListByPost:        query.NewBookingListByPostHandler(bookingRepo, cnf.CacheSrv),
			BookingListByUser:        query.NewBookingListByUserHandler(bookingRepo, cnf.CacheSrv),
			BookingListMyAttendees:   query.NewBookingListMyAttendeesHandler(bookingRepo),
			BookingListMyOrganized:   query.NewBookingListMyOrganizedHandler(bookingRepo),
			BookingView:              query.NewBookingViewHandler(bookingRepo),

			InviteGetByBookingUUID: query.NewInviteGetByBookingUUIDHandler(inviteRepo),
			InviteGetByEmail:       query.NewInviteGetByEmailHandler(inviteRepo),
			InviteGetByUUID:        query.NewInviteGetByUUIDHandler(inviteRepo),
		},
	}
}
