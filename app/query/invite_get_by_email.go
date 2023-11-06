package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/invite"
)

type InviteGetByEmailQuery struct {
	UserEmail string
}

type InviteGetByEmailRes struct {
	Invites []*invite.Entity
}

type InviteGetByEmailHandler cqrs.HandlerFunc[InviteGetByEmailQuery, *InviteGetByEmailRes]

func NewInviteGetByEmailHandler(repo invite.Repository) InviteGetByEmailHandler {
	return func(ctx context.Context, query InviteGetByEmailQuery) (*InviteGetByEmailRes, *i18np.Error) {
		res, err := repo.GetByEmail(ctx, query.UserEmail)
		if err != nil {
			return nil, err
		}
		return &InviteGetByEmailRes{Invites: res}, nil
	}
}
