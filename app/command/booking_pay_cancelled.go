package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingPayCancelledCmd struct {
	BookingUUID string `json:"bookingUUID"`
}

type BookingPayCancelledRes struct{}

type BookingPayCancelledHandler cqrs.HandlerFunc[BookingPayCancelledCmd, *BookingPayCancelledRes]

func NewBookingPayCancelledHandler(repo booking.Repo) BookingPayCancelledHandler {
	return func(ctx context.Context, cmd BookingPayCancelledCmd) (*BookingPayCancelledRes, *i18np.Error) {
		err := repo.MarkPayCancelled(ctx, cmd.BookingUUID)
		if err != nil {
			return nil, err
		}
		return &BookingPayCancelledRes{}, nil
	}
}
