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
	BookingPayCancelled      command.BookingPayCancelledHandler
	BookingPaySuccess        command.BookingPaySuccessHandler
	BookingMarkPublic        command.BookingMarkPublicHandler
	BookingRemoveGuest       command.BookingRemoveGuestHandler
	BookingGuestMarkPublic   command.BookingGuestMarkPublicHandler
	BookingGuestMarkPrivate  command.BookingGuestMarkPrivateHandler
	BookingValidationSucceed command.BookingValidationSucceedHandler
	BookingValidationFailed  command.BookingValidationFailedHandler
	BookingCancelAsAdmin     command.BookingCancelAsAdminHandler
	BookingCancelAsBusiness  command.BookingCancelAsBusinessHandler

	InviteCreate command.InviteCreateHandler
	InviteUse    command.InviteUseHandler
	InviteDelete command.InviteDeleteHandler
}

type Queries struct {
	BookingAdminList         query.BookingAdminListHandler
	BookingAdminView         query.BookingAdminViewHandler
	BookingCheckAvailability query.BookingCheckAvailabilityHandler
	BookingListByBusiness    query.BookingListByBusinessHandler
	BookingListByListing     query.BookingListByListingHandler
	BookingListByUser        query.BookingListByUserHandler
	BookingList              query.BookingListHandler
	BookingView              query.BookingViewHandler
	BookingViewBusiness      query.BookingViewBusinessHandler

	InviteGetByBookingUUID query.InviteGetByBookingUUIDHandler
	InviteGetByEmail       query.InviteGetByEmailHandler
	InviteGetByUUID        query.InviteGetByUUIDHandler
}
