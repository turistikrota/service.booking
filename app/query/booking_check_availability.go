package query

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingCheckAvailabilityQuery struct {
	ListingUUID string `query:"-"`
	StartDate   string `query:"start" validate:"required,datetime=2006-01-02"`
	EndDate     string `query:"end" validate:"required,datetime=2006-01-02"`
}

type BookingCheckAvailabilityRes struct {
	IsAvailable bool
}

type BookingCheckAvailabilityHandler cqrs.HandlerFunc[BookingCheckAvailabilityQuery, *BookingCheckAvailabilityRes]

func NewBookingCheckAvailabilityHandler(factory booking.Factory, repo booking.Repo) BookingCheckAvailabilityHandler {
	return func(ctx context.Context, query BookingCheckAvailabilityQuery) (*BookingCheckAvailabilityRes, *i18np.Error) {
		startDate, _ := time.Parse("2006-01-02", query.StartDate)
		endDate, _ := time.Parse("2006-01-02", query.EndDate)
		dateError := factory.ValidateDateTime(startDate, endDate)
		if dateError != nil {
			return nil, dateError
		}
		res, err := repo.CheckAvailability(ctx, query.ListingUUID, startDate, endDate)
		if err != nil {
			return nil, err
		}
		return &BookingCheckAvailabilityRes{IsAvailable: res}, nil
	}
}
