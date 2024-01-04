package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/domains/listing"
)

type BookingValidationSucceedCmd struct {
	BookingUUID  string                `json:"booking_uuid"`
	ListingUUID  string                `json:"listing_uuid"`
	BusinessUUID string                `json:"business_uuid"`
	BusinessName string                `json:"business_name"`
	TotalPrice   float64               `json:"total_price"`
	PricePerDays []listing.PricePerDay `json:"price_per_days"`
}

type BookingValidationSucceedRes struct{}

type BookingValidationSucceedHandler cqrs.HandlerFunc[BookingValidationSucceedCmd, *BookingValidationSucceedRes]

func NewBookingValidationSucceedHandler(repo booking.Repo, events booking.Events) BookingValidationSucceedHandler {
	return func(ctx context.Context, cmd BookingValidationSucceedCmd) (*BookingValidationSucceedRes, *i18np.Error) {
		book, err := repo.GetByUUID(ctx, cmd.BookingUUID)
		if err != nil {
			return nil, err
		}
		days := make([]booking.Day, len(cmd.PricePerDays))
		for i, pricePerDay := range cmd.PricePerDays {
			days[i] = booking.Day{
				Date:  pricePerDay.Date,
				Price: pricePerDay.Price,
			}
		}
		_err := repo.Validated(ctx, &booking.Validated{
			UUID:         cmd.BookingUUID,
			ListingUUID:  cmd.ListingUUID,
			BusinessUUID: cmd.BusinessUUID,
			BusinessName: cmd.BusinessName,
			TotalPrice:   cmd.TotalPrice,
			Days:         days,
		})
		if _err != nil {
			return nil, _err
		}
		events.PayPending(booking.PayPendingEvent{
			BookingUUID:  cmd.BookingUUID,
			BusinessUUID: cmd.BusinessUUID,
			ListingUUID:  cmd.ListingUUID,
			Price:        cmd.TotalPrice,
			User:         &book.User,
		})
		return &BookingValidationSucceedRes{}, nil
	}
}
