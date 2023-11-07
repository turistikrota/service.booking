package command

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/domains/invite"
)

type InviteCreateCmd struct {
	UserUUID    string `json:"-"`
	UserName    string `json:"-"`
	BookingUUID string `json:"-"`
	Email       string `json:"email" validate:"required,email"`
	Locale      string `json:"locale" validate:"required,locale"`
}

type InviteCreateRes struct{}

type InviteCreateHandler cqrs.HandlerFunc[InviteCreateCmd, *InviteCreateRes]

func NewInviteCreateHandler(bookingFactory booking.Factory, bookingRepo booking.Repo, factory invite.Factory, repo invite.Repository, events invite.Events) InviteCreateHandler {
	return func(ctx context.Context, cmd InviteCreateCmd) (*InviteCreateRes, *i18np.Error) {
		fmt.Println("check 1")
		_, exists, err := bookingRepo.GetDetailWithUser(ctx, cmd.BookingUUID, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		fmt.Println("check 2")
		if exists != nil && !*exists {
			return nil, bookingFactory.Errors.OnlyAdminCanDoThisAction()
		}
		fmt.Println("check 3")
		res, _err := repo.Create(ctx, factory.New(cmd.Email, cmd.BookingUUID, cmd.UserName))
		if _err != nil {
			return nil, _err
		}
		fmt.Println("check 4")
		events.Invite(invite.InviteEvent{
			Locale:     cmd.Locale,
			Email:      cmd.Email,
			InviteUUID: res.UUID,
			UserName:   cmd.UserName,
		})
		fmt.Println("check 5")
		return &InviteCreateRes{}, nil
	}
}
