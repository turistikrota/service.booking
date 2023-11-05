package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingValidationFailedCmd struct {
	BookingUUID string                     `json:"booking_uuid"`
	PostUUID    string                     `json:"post_uuid"`
	Errors      []*booking.ValidationError `json:"errors"`
}

type BookingValidationFailedRes struct{}

type BookingValidationFailedHandler cqrs.HandlerFunc[BookingValidationFailedCmd, *BookingValidationFailedRes]

func NewBookingValidationFailedHandler(repo booking.Repo) BookingValidationFailedHandler {
	return func(ctx context.Context, cmd BookingValidationFailedCmd) (*BookingValidationFailedRes, *i18np.Error) {
		err := repo.MarkNotValid(ctx, cmd.BookingUUID)
		if err != nil {
			return nil, err
		}
		return &BookingValidationFailedRes{}, nil
	}
}
