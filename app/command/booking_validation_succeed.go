package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/domains/post"
)

type BookingValidationSucceedCmd struct {
	BookingUUID  string             `json:"booking_uuid"`
	PostUUID     string             `json:"post_uuid"`
	OwnerUUID    string             `json:"owner_uuid"`
	OwnerName    string             `json:"owner_name"`
	TotalPrice   float64            `json:"total_price"`
	PricePerDays []post.PricePerDay `json:"price_per_days"`
}

type BookingValidationSucceedRes struct{}

type BookingValidationSucceedHandler cqrs.HandlerFunc[BookingValidationSucceedCmd, *BookingValidationSucceedRes]

func NewBookingValidationSucceedHandler(repo booking.Repo, events booking.Events) BookingValidationSucceedHandler {
	return func(ctx context.Context, cmd BookingValidationSucceedCmd) (*BookingValidationSucceedRes, *i18np.Error) {
		days := make([]booking.Day, len(cmd.PricePerDays))
		for i, pricePerDay := range cmd.PricePerDays {
			days[i] = booking.Day{
				Date:  pricePerDay.Date,
				Price: pricePerDay.Price,
			}
		}
		_err := repo.Validated(ctx, &booking.Validated{
			UUID:       cmd.BookingUUID,
			PostUUID:   cmd.PostUUID,
			OwnerUUID:  cmd.OwnerUUID,
			OwnerName:  cmd.OwnerName,
			TotalPrice: cmd.TotalPrice,
			Days:       days,
		})
		if _err != nil {
			return nil, _err
		}
		events.PayPending(booking.PayPendingEvent{
			BookingUUID: cmd.BookingUUID,
		})
		return &BookingValidationSucceedRes{}, nil
	}
}
