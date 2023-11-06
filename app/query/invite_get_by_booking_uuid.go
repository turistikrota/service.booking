package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/invite"
)

type InviteGetByBookingUUIDQuery struct {
	BookingUUID string `params:"uuid" validate:"required,object_id"`
}

type InviteGetByBookingUUIDRes struct {
	Invites []*invite.Entity
}

type InviteGetByBookingUUIDHandler cqrs.HandlerFunc[InviteGetByBookingUUIDQuery, *InviteGetByBookingUUIDRes]

func NewInviteGetByBookingUUIDHandler(repo invite.Repository) InviteGetByBookingUUIDHandler {
	return func(ctx context.Context, query InviteGetByBookingUUIDQuery) (*InviteGetByBookingUUIDRes, *i18np.Error) {
		res, err := repo.GetByBookingUUID(ctx, query.BookingUUID)
		if err != nil {
			return nil, err
		}
		return &InviteGetByBookingUUIDRes{Invites: res}, nil
	}
}
