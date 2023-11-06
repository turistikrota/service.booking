package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/invite"
)

type InviteGetByUUIDQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type InviteGetByUUIDRes struct {
	Invite *invite.Entity
}

type InviteGetByUUIDHandler cqrs.HandlerFunc[InviteGetByUUIDQuery, *InviteGetByUUIDRes]

func NewInviteGetByUUIDHandler(repo invite.Repository) InviteGetByUUIDHandler {
	return func(ctx context.Context, query InviteGetByUUIDQuery) (*InviteGetByUUIDRes, *i18np.Error) {
		res, err := repo.GetByUUID(ctx, query.UUID)
		if err != nil {
			return nil, err
		}
		return &InviteGetByUUIDRes{Invite: res}, nil
	}
}
