package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingCreateCmd struct{}

type BookingCreateRes struct{}

type BookingCreateHandler cqrs.HandlerFunc[BookingCreateCmd, *BookingCreateRes]

func NewBookingCreateHandler(factory booking.Factory, repo booking.Repo, events booking.Events) BookingCreateHandler {
	return func(ctx context.Context, cmd BookingCreateCmd) (*BookingCreateRes, *i18np.Error) {
		return &BookingCreateRes{}, nil
	}
}
