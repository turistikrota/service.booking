package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingCancelAsAdminContentDto struct {
	Content string `json:"content" validate:"required,max=1000"`
}

type BookingCancelAsAdminCmd struct {
	UUID      string                         `json:"-"`
	TrContent BookingCancelAsAdminContentDto `json:"tr" validate:"required,dive"`
	EnContent BookingCancelAsAdminContentDto `json:"en" validate:"required,dive"`
}

type BookingCancelAsAdminRes struct{}

type BookingCancelAsAdminHandler cqrs.HandlerFunc[BookingCancelAsAdminCmd, *BookingCancelAsAdminRes]

func NewBookingCancelAsAdminHandler(factory booking.Factory, repo booking.Repo, events booking.Events) BookingCancelAsAdminHandler {
	return func(ctx context.Context, cmd BookingCancelAsAdminCmd) (*BookingCancelAsAdminRes, *i18np.Error) {
		if err := repo.CancelAsAdmin(ctx, cmd.UUID, factory.NewCancelReason(booking.NewCancelConfig{
			TrContent: cmd.TrContent.Content,
			EnContent: cmd.EnContent.Content,
			IsAdmin:   true,
		})); err != nil {
			return nil, err
		}
		events.Cancelled(booking.CancelledEvent{
			BookingUUID: cmd.UUID,
		})
		return &BookingCancelAsAdminRes{}, nil
	}
}
