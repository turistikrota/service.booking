package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingPayTimeoutCmd struct {
	BookingUUID string `json:"bookingUUID"`
}

type BookingPayTimeoutRes struct{}

type BookingPayTimeoutHandler cqrs.HandlerFunc[BookingPayTimeoutCmd, *BookingPayTimeoutRes]

func NewBookingPayTimeoutHandler(repo booking.Repo) BookingPayTimeoutHandler {
	return func(ctx context.Context, cmd BookingPayTimeoutCmd) (*BookingPayTimeoutRes, *i18np.Error) {
		err := repo.MarkExpired(ctx, cmd.BookingUUID)
		if err != nil {
			return nil, err
		}
		return &BookingPayTimeoutRes{}, nil
	}
}
