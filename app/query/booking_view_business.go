package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingViewBusinessQuery struct {
	UUID         string `params:"uuid" validate:"required,object_id"`
	BusinessUUID string `params:"-" query:"-"`
}

type BookingViewBusinessRes struct {
	Detail booking.BookingBusinessViewDto
}

type BookingViewBusinessHandler cqrs.HandlerFunc[BookingViewBusinessQuery, *BookingViewBusinessRes]

func NewBookingViewBusinessHandler(repo booking.Repo) BookingViewBusinessHandler {
	return func(ctx context.Context, query BookingViewBusinessQuery) (*BookingViewBusinessRes, *i18np.Error) {
		res, err := repo.GetByUUIDAsBusiness(ctx, query.UUID, booking.WithBusiness{
			UUID: query.BusinessUUID,
		})
		if err != nil {
			return nil, err
		}
		return &BookingViewBusinessRes{
			Detail: res.ToBookingBusinessViewDto(),
		}, nil
	}
}
