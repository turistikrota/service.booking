package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingMarkPrivateCmd struct {
	UserUUID string `params:"-"`
	UUID     string `params:"uuid" validate:"required,object_id"`
}

type BookingMarkPrivateRes struct{}

type BookingMarkPrivateHandler cqrs.HandlerFunc[BookingMarkPrivateCmd, *BookingMarkPrivateRes]

func NewBookingMarkPrivateHandler(factory booking.Factory, repo booking.Repo) BookingMarkPrivateHandler {
	return func(ctx context.Context, cmd BookingMarkPrivateCmd) (*BookingMarkPrivateRes, *i18np.Error) {
		_, exists, err := repo.GetDetailWithUser(ctx, cmd.UUID, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		if exists != nil && !*exists {
			return nil, factory.Errors.OnlyAdminCanDoThisAction()
		}
		if err := repo.MarkPrivate(ctx, cmd.UUID); err != nil {
			return nil, err
		}
		return &BookingMarkPrivateRes{}, nil
	}
}
