package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingPaySuccessCmd struct {
	BookingUUID string `json:"bookingUUID"`
}

type BookingPaySuccessRes struct{}

type BookingPaySuccessHandler cqrs.HandlerFunc[BookingPaySuccessCmd, *BookingPaySuccessRes]

func NewBookingPaySuccessHandler(repo booking.Repo) BookingPaySuccessHandler {
	return func(ctx context.Context, cmd BookingPaySuccessCmd) (*BookingPaySuccessRes, *i18np.Error) {
		err := repo.MarkPaid(ctx, cmd.BookingUUID)
		if err != nil {
			return nil, err
		}
		return &BookingPaySuccessRes{}, nil
	}
}
