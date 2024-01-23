package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingCancelAsBusinessContentDto struct {
	Content string `json:"content" validate:"required,max=1000"`
}

type BookingCancelAsBusinessCmd struct {
	UUID         string                            `json:"-"`
	BusinessUUID string                            `json:"-"`
	TrContent    BookingCancelAsBusinessContentDto `json:"trContent" validate:"required,dive"`
	EnContent    BookingCancelAsBusinessContentDto `json:"enContent" validate:"required,dive"`
}

type BookingCancelAsBusinessRes struct{}

type BookingCancelAsBusinessHandler cqrs.HandlerFunc[BookingCancelAsBusinessCmd, *BookingCancelAsBusinessRes]

func NewBookingCancelAsBusinessHandler(factory booking.Factory, repo booking.Repo, events booking.Events) BookingCancelAsBusinessHandler {
	return func(ctx context.Context, cmd BookingCancelAsBusinessCmd) (*BookingCancelAsBusinessRes, *i18np.Error) {
		bus := booking.WithBusiness{
			UUID: cmd.BusinessUUID,
		}
		book, err := repo.GetByUUIDAsBusiness(ctx, cmd.UUID, bus)
		if err != nil {
			return nil, err
		}
		if cancellable := factory.IsCancelableAsBusiness(book); !cancellable {
			return nil, factory.Errors.NotCancelable()
		}
		if err := repo.CancelAsBusiness(ctx, cmd.UUID, bus, factory.NewCancelReason(booking.NewCancelConfig{
			TrContent: cmd.TrContent.Content,
			EnContent: cmd.EnContent.Content,
			IsAdmin:   false,
		})); err != nil {
			return nil, err
		}
		events.Cancelled(booking.CancelledEvent{
			BookingUUID: cmd.UUID,
		})
		return &BookingCancelAsBusinessRes{}, nil
	}
}
