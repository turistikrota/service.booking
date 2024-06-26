package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.booking/app/command"
	"github.com/turistikrota/service.booking/app/query"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/pkg/utils"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_business"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h srv) BookingCreate(ctx *fiber.Ctx) error {
	detail := command.ListingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.BookingCreateCmd{}
	cmd.ListingUUID = detail.ListingUUID
	cmd.User = booking.User{
		UUID: current_user.Parse(ctx).UUID,
		Name: current_account.Parse(ctx).Name,
	}
	h.parseBody(ctx, &cmd)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	cmd.Locale = l
	res, err := h.app.Commands.BookingCreate(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingCancel(ctx *fiber.Ctx) error {
	cmd := command.BookingCancelCmd{}
	h.parseParams(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.BookingCancel(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingCancelAsAdmin(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.BookingCancelAsAdminCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.BookingCancelAsAdmin(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingCancelAsBusiness(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.BookingCancelAsBusinessCmd{}
	h.parseBody(ctx, &cmd)
	cmd.BusinessUUID = current_business.Parse(ctx).UUID
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.BookingCancelAsBusiness(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingMarkPrivate(ctx *fiber.Ctx) error {
	cmd := command.BookingMarkPrivateCmd{}
	h.parseParams(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.BookingMarkPrivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingMarkPublic(ctx *fiber.Ctx) error {
	cmd := command.BookingMarkPublicCmd{}
	h.parseParams(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.BookingMarkPublic(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingGuestRemove(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.BookingRemoveGuestCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.BookingRemoveGuest(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingGuestMarkPublic(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.BookingGuestMarkPublicCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.BookingGuestMarkPublic(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingGuestMarkPrivate(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.BookingGuestMarkPrivateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.BookingGuestMarkPrivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) InviteCreate(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.InviteCreateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.BookingUUID = detail.UUID
	res, err := h.app.Commands.InviteCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) InviteUse(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.InviteUseCmd{}
	h.parseBody(ctx, &cmd)
	u := current_user.Parse(ctx)
	cmd.UserUUID = u.UUID
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserEmail = u.Email
	cmd.InviteUUID = detail.UUID
	res, err := h.app.Commands.InviteUse(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) InviteDelete(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.InviteDeleteCmd{}
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.InviteDelete(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingAdminList(ctx *fiber.Ctx) error {
	p := utils.Pagination{}
	h.parseQuery(ctx, &p)
	filter := booking.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.BookingAdminListQuery{}
	query.Pagination = &p
	query.FilterEntity = &filter
	res, err := h.app.Queries.BookingAdminList(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) BookingList(ctx *fiber.Ctx) error {
	p := utils.Pagination{}
	h.parseQuery(ctx, &p)
	filter := booking.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.BookingListQuery{}
	query.Pagination = &p
	query.FilterEntity = &filter
	query.UserUUID = current_user.Parse(ctx).UUID
	query.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Queries.BookingList(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) BookingAdminView(ctx *fiber.Ctx) error {
	query := query.BookingAdminViewQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.BookingAdminView(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) BookingCheckAvailability(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	query := query.BookingCheckAvailabilityQuery{}
	h.parseQuery(ctx, &query)
	query.ListingUUID = detail.UUID
	res, err := h.app.Queries.BookingCheckAvailability(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.IsAvailable)
}

func (h srv) BookingListByBusiness(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	p := utils.Pagination{}
	h.parseQuery(ctx, &p)
	filter := booking.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.BookingListByBusinessQuery{}
	query.Pagination = &p
	query.FilterEntity = &filter
	query.BusinessUUID = detail.UUID
	query.IsPublic = true
	res, err := h.app.Queries.BookingListByBusiness(ctx.UserContext(), query)
	if err != nil {
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) BookingListByBusinessAuthorized(ctx *fiber.Ctx) error {
	p := utils.Pagination{}
	h.parseQuery(ctx, &p)
	filter := booking.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.BookingListByBusinessQuery{}
	query.Pagination = &p
	query.FilterEntity = &filter
	query.BusinessUUID = current_business.Parse(ctx).UUID
	query.IsPublic = false
	res, err := h.app.Queries.BookingListByBusiness(ctx.UserContext(), query)
	if err != nil {
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) BookingViewByBusiness(ctx *fiber.Ctx) error {
	query := query.BookingViewBusinessQuery{}
	h.parseParams(ctx, &query)
	query.BusinessUUID = current_business.Parse(ctx).UUID
	res, err := h.app.Queries.BookingViewBusiness(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) BookingListByListing(ctx *fiber.Ctx) error {
	detail := command.BookingDetailCmd{}
	h.parseParams(ctx, &detail)
	p := utils.Pagination{}
	h.parseQuery(ctx, &p)
	filter := booking.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.BookingListByListingQuery{}
	query.Pagination = &p
	query.FilterEntity = &filter
	query.ListingUUID = detail.UUID
	query.IsPublic = true
	res, err := h.app.Queries.BookingListByListing(ctx.UserContext(), query)
	if err != nil {
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) BookingListByUser(ctx *fiber.Ctx) error {
	p := utils.Pagination{}
	h.parseQuery(ctx, &p)
	query := query.BookingListByUserQuery{}
	h.parseParams(ctx, &query)
	filter := booking.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query.Pagination = &p
	query.FilterEntity = &filter
	res, err := h.app.Queries.BookingListByUser(ctx.UserContext(), query)
	if err != nil {
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) BookingView(ctx *fiber.Ctx) error {
	query := query.BookingViewQuery{}
	h.parseParams(ctx, &query)
	account := current_account.Parse(ctx)
	if account != nil {
		query.UserName = account.Name
		query.UserId = current_user.Parse(ctx).UUID
	}
	res, err := h.app.Queries.BookingView(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) InviteGetByBookingUUID(ctx *fiber.Ctx) error {
	query := query.InviteGetByBookingUUIDQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.InviteGetByBookingUUID(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Invites)
}

func (h srv) InviteGetByEmail(ctx *fiber.Ctx) error {
	query := query.InviteGetByEmailQuery{}
	query.UserEmail = current_user.Parse(ctx).Email
	res, err := h.app.Queries.InviteGetByEmail(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Invites)
}

func (h srv) InviteGetByUUID(ctx *fiber.Ctx) error {
	query := query.InviteGetByUUIDQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.InviteGetByUUID(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Invite)
}
