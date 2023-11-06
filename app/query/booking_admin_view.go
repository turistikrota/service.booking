package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingAdminViewQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type BookingAdminViewRes struct {
	Res booking.BookingAdminViewDto
}

type BookingAdminViewHandler cqrs.HandlerFunc[BookingAdminViewQuery, *BookingAdminViewRes]

func NewBookingAdminViewHandler(repo booking.Repo) BookingAdminViewHandler {
	return func(ctx context.Context, query BookingAdminViewQuery) (*BookingAdminViewRes, *i18np.Error) {
		res, err := repo.GetByUUID(ctx, query.UUID)
		if err != nil {
			return nil, err
		}
		return &BookingAdminViewRes{
			Res: res.ToAdminViewDto(),
		}, nil
	}
}
