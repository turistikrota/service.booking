package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingViewQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type BookingViewRes struct {
	Detail booking.BookingViewDto
}

type BookingViewHandler cqrs.HandlerFunc[BookingViewQuery, *BookingViewRes]

func NewBookingViewHandler(repo booking.Repo) BookingViewHandler {
	return func(ctx context.Context, query BookingViewQuery) (*BookingViewRes, *i18np.Error) {
		res, err := repo.GetByUUID(ctx, query.UUID)
		if err != nil {
			return nil, err
		}
		return &BookingViewRes{
			Detail: res.ToViewDto(),
		}, nil
	}
}
