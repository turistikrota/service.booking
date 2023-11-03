package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/domains/booking"
)

type BookingCreateCmd struct {
	PostUUID  string          `json:"-"`
	User      booking.User    `json:"-"`
	People    *booking.People `json:"people" validate:"required"`
	StartDate string          `json:"startDate" validate:"required,datetime"`
	EndDate   string          `json:"endDate" validate:"required,datetime"`
	IsPublic  *bool           `json:"isPublic" validate:"required"`
}

type BookingCreateRes struct {
	UUID string `json:"uuid"`
}

type BookingCreateHandler cqrs.HandlerFunc[BookingCreateCmd, *BookingCreateRes]

func NewBookingCreateHandler(factory booking.Factory, repo booking.Repo, events booking.Events) BookingCreateHandler {
	return func(ctx context.Context, cmd BookingCreateCmd) (*BookingCreateRes, *i18np.Error) {
		startDate, _ := time.Parse("2006-01-02", cmd.StartDate)
		endDate, _ := time.Parse("2006-01-02", cmd.EndDate)
		e := factory.New(booking.NewConfig{
			PostUUID:  cmd.PostUUID,
			People:    *cmd.People,
			User:      cmd.User,
			State:     booking.Created,
			StartDate: startDate,
			EndDate:   endDate,
			IsPublic:  cmd.IsPublic,
		})
		res, err := repo.Create(ctx, e)
		if err != nil {
			return nil, err
		}
		return &BookingCreateRes{
			UUID: res.UUID,
		}, nil
	}
}
