package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingCancelCmd struct {
	UserUUID string `params:"-"`
	UUID     string `params:"uuid" validate:"required,object_id"`
}

type BookingCancelRes struct{}

type BookingCancelHandler cqrs.HandlerFunc[BookingCancelCmd, *BookingCancelRes]

func NewBookingCancelHandler(factory booking.Factory, repo booking.Repo, events booking.Events) BookingCancelHandler {
	return func(ctx context.Context, cmd BookingCancelCmd) (*BookingCancelRes, *i18np.Error) {
		book, exists, err := repo.GetDetailWithUser(ctx, cmd.UUID, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		if exists != nil && !*exists {
			return nil, factory.Errors.OnlyAdminCanDoThisAction()
		}
		if cancellable := factory.IsCancelable(book); !cancellable {
			return nil, factory.Errors.NotCancelable()
		}
		if err := repo.Cancel(ctx, cmd.UUID); err != nil {
			return nil, err
		}
		events.Cancelled(booking.CancelledEvent{
			BookingUUID: cmd.UUID,
		})
		return &BookingCancelRes{}, nil
	}
}
