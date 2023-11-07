package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.booking/app/command"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h srv) BookingCreate(ctx *fiber.Ctx) error {
	detail := command.PostDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.BookingCreateCmd{}
	cmd.PostUUID = detail.PostUUID
	cmd.User = booking.User{
		UUID: current_user.Parse(ctx).UUID,
		Name: current_account.Parse(ctx).Name,
	}
	h.parseBody(ctx, &cmd)
	res, err := h.app.Commands.BookingCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BookingCancel(ctx *fiber.Ctx) error {
	cmd := command.BookingCancelCmd{}
	h.parseParams(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.BookingCancel(ctx.UserContext(), cmd)
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
