package app

import (
	"github.com/turistikrota/service.booking/app/command"
	"github.com/turistikrota/service.booking/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	BookingCreate            command.BookingCreateHandler
	BookingCancel            command.BookingCancelHandler
	BookingMarkPrivate       command.BookingMarkPrivateHandler
	BookingPayTimeout        command.BookingPayTimeoutHandler
	BookingPaySuccess        command.BookingPaySuccessHandler
	BookingMarkPublic        command.BookingMarkPublicHandler
	BookingRemoveGuest       command.BookingRemoveGuestHandler
	BookingGuestMarkPublic   command.BookingGuestMarkPublicHandler
	BookingGuestMarkPrivate  command.BookingGuestMarkPrivateCmd
	BookingValidationSucceed command.BookingValidationSucceedHandler
	BookingValidationFailed  command.BookingValidationFailedHandler

	InviteCreate command.InviteCreateHandler
	InviteUse    command.InviteUseHandler
	InviteDelete command.InviteDeleteHandler
}

type Queries struct {
	InviteGetByBookingUUID   query.InviteGetByBookingUUIDHandler
	InviteGetByEmail         query.InviteGetByEmailHandler
	InviteGetByUUID          query.InviteGetByUUIDHandler
	BookingAdminList         query.BookingAdminListHandler
	BookingAdminView         query.BookingAdminViewHandler
	BookingCheckAvailability query.BookingCheckAvailabilityHandler
	BookingListByOwner       query.BookingListByOwnerHandler
	BookingListByPost        query.BookingListByPostHandler
	BookingListByUser        query.BookingListByUserHandler
	BookingListMyAttendees   query.BookingListMyAttendeesHandler
	BookingListMyOrganized   query.BookingListMyOrganizedHandler
	BookingView              query.BookingViewHandler
}
