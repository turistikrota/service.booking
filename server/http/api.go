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
