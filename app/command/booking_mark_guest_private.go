package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingGuestMarkPrivateCmd struct {
	UserUUID  string `json:"-"`
	UserName  string `json:"-"`
	UUID      string `json:"-"`
	GuestName string `json:"guestName" validate:"required"`
	GuestUUID string `json:"guestUUID" validate:"required,object_id"`
}

type BookingGuestMarkPrivateRes struct{}

type BookingGuestMarkPrivateHandler cqrs.HandlerFunc[BookingGuestMarkPrivateCmd, *BookingGuestMarkPrivateRes]

func NewBookingGuestMarkPrivateHandler(factory booking.Factory, repo booking.Repo, events booking.Events) BookingGuestMarkPrivateHandler {
	return func(ctx context.Context, cmd BookingGuestMarkPrivateCmd) (*BookingGuestMarkPrivateRes, *i18np.Error) {
		_, exists, err := repo.GetDetailWithUser(ctx, cmd.UUID, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		if exists != nil && !*exists {
			return nil, factory.Errors.OnlyAdminCanDoThisAction()
		}
		if err := repo.MarkGuestAsPrivate(ctx, cmd.UUID, booking.WithUser{
			UUID: cmd.GuestUUID,
			Name: cmd.GuestName,
		}, booking.WithUser{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		}); err != nil {
			return nil, err
		}
		return &BookingGuestMarkPrivateRes{}, nil
	}
}
