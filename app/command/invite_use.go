package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/domains/invite"
)

type InviteUseCmd struct {
	UserUUID   string `json:"-"`
	UserName   string `json:"-"`
	UserEmail  string `json:"-"`
	InviteUUID string `json:"-"`
	IsPublic   *bool  `json:"isPublic" validate:"required"`
}

type InviteUseRes struct{}

type InviteUseHandler cqrs.HandlerFunc[InviteUseCmd, *InviteUseRes]

func NewInviteUseHandler(factory invite.Factory, repo invite.Repository, bookingRepo booking.Repo) InviteUseHandler {
	return func(ctx context.Context, cmd InviteUseCmd) (*InviteUseRes, *i18np.Error) {
		res, err := repo.GetByUUID(ctx, cmd.InviteUUID)
		if err != nil {
			return nil, err
		}
		if res.Email != cmd.UserEmail {
			return nil, factory.Errors.EmailMismatch()
		}
		if res.IsUsed {
			return nil, factory.Errors.Used()
		}
		if res.IsDeleted {
			return nil, factory.Errors.Deleted()
		}
		if res.CreatedAt.Add(24 * time.Hour).Before(time.Now()) {
			return nil, factory.Errors.Timeout()
		}
		if res.CreatorUserName == cmd.UserName {
			return nil, factory.Errors.SameUser()
		}
		err = bookingRepo.AddGuest(ctx, res.BookingUUID, &booking.Guest{
			UUID:     cmd.UserUUID,
			Name:     cmd.UserName,
			IsPublic: *cmd.IsPublic,
		})
		if err != nil {
			return nil, err
		}
		err = repo.Use(ctx, cmd.InviteUUID)
		if err != nil {
			return nil, err
		}
		return &InviteUseRes{}, nil
	}
}
