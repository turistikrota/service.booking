package query

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingCheckAvailabilityQuery struct {
	PostUUID  string `query:"-"`
	StartDate string `query:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate   string `query:"end_date" validate:"required,datetime=2006-01-02"`
}

type BookingCheckAvailabilityRes struct {
	IsAvailable bool
}

type BookingCheckAvailabilityHandler cqrs.HandlerFunc[BookingCheckAvailabilityQuery, *BookingCheckAvailabilityRes]

func NewBookingCheckAvailabilityHandler(repo booking.Repo) BookingCheckAvailabilityHandler {
	return func(ctx context.Context, query BookingCheckAvailabilityQuery) (*BookingCheckAvailabilityRes, *i18np.Error) {
		startDate, _ := time.Parse("2006-01-02", query.StartDate)
		endDate, _ := time.Parse("2006-01-02", query.EndDate)
		res, err := repo.CheckAvailability(ctx, query.PostUUID, startDate, endDate)
		if err != nil {
			return nil, err
		}
		return &BookingCheckAvailabilityRes{IsAvailable: res}, nil
	}
}
