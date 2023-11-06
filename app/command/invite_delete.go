package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/invite"
)

type InviteDeleteCmd struct {
	UserName string `params:"-"`
	UUID     string `params:"uuid" validate:"required,object_id"`
}

type InviteDeleteRes struct{}

type InviteDeleteHandler cqrs.HandlerFunc[InviteDeleteCmd, *InviteDeleteRes]

func NewInviteDeleteHandler(repo invite.Repository) InviteDeleteHandler {
	return func(ctx context.Context, cmd InviteDeleteCmd) (*InviteDeleteRes, *i18np.Error) {
		err := repo.Delete(ctx, cmd.UUID, cmd.UserName)
		if err != nil {
			return nil, err
		}
		return &InviteDeleteRes{}, nil
	}
}
