package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingMarkPublicCmd struct {
	UserUUID string `params:"-"`
	UserName string `params:"-"`
	UUID     string `params:"uuid" validate:"required,object_id"`
}

type BookingMarkPublicRes struct{}

type BookingMarkPublicHandler cqrs.HandlerFunc[BookingMarkPublicCmd, *BookingMarkPublicRes]

func NewBookingMarkPublicHandler(factory booking.Factory, repo booking.Repo) BookingMarkPublicHandler {
	return func(ctx context.Context, cmd BookingMarkPublicCmd) (*BookingMarkPublicRes, *i18np.Error) {
		_, exists, err := repo.GetDetailWithUser(ctx, cmd.UUID, cmd.UserUUID, cmd.UserName)
		if err != nil {
			return nil, err
		}
		if exists != nil && !*exists {
			return nil, factory.Errors.OnlyAdminCanDoThisAction()
		}
		if err := repo.MarkPublic(ctx, cmd.UUID); err != nil {
			return nil, err
		}
		return &BookingMarkPublicRes{}, nil
	}
}
