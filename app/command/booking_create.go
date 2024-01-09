package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.booking/config"
	"github.com/turistikrota/service.booking/domains/booking"
	listing "github.com/turistikrota/service.booking/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingCreateCmd struct {
	Locale      string          `json:"-"`
	ListingUUID string          `json:"-"`
	User        booking.User    `json:"-"`
	People      *booking.People `json:"people" validate:"required"`
	StartDate   string          `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate     string          `json:"endDate" validate:"required,datetime=2006-01-02"`
	IsPublic    *bool           `json:"isPublic" validate:"required"`
}

type BookingCreateRes struct {
	UUID string `json:"uuid"`
}

type BookingCreateHandler cqrs.HandlerFunc[BookingCreateCmd, *BookingCreateRes]

func NewBookingCreateHandler(factory booking.Factory, repo booking.Repo, events booking.Events, rpcConfig config.Rpc) BookingCreateHandler {

	getListing := func(ctx context.Context, uuid string, locale string) (*listing.Entity, error) {
		var opt grpc.DialOption
		if !rpcConfig.ListingUsesSsl {
			opt = grpc.WithTransportCredentials(insecure.NewCredentials())
		} else {
			opt = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
		}
		conn, err := grpc.Dial(rpcConfig.ListingHost, opt)
		if err != nil {
			return nil, err
		}
		defer conn.Close()
		c := listing.NewListingServiceClient(conn)
		response, err := c.GetEntity(ctx, &listing.GetEntityRequest{
			Uuid:   uuid,
			Locale: locale,
		})
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	return func(ctx context.Context, cmd BookingCreateCmd) (*BookingCreateRes, *i18np.Error) {
		startDate, _ := time.Parse("2006-01-02", cmd.StartDate)
		endDate, _ := time.Parse("2006-01-02", cmd.EndDate)
		available, err := repo.CheckAvailability(ctx, cmd.ListingUUID, startDate, endDate)
		if err != nil {
			return nil, err
		}
		if !available {
			return nil, factory.Errors.NotAvailable()
		}
		listing, _err := getListing(ctx, cmd.ListingUUID, cmd.Locale)
		if _err != nil {
			return nil, factory.Errors.Failed(_err.Error())
		}
		listingImages := make([]booking.ListingImage, len(listing.Images))
		for i, image := range listing.Images {
			listingImages[i] = booking.ListingImage{
				Url:   image.Url,
				Order: int(image.Order),
			}
		}
		e := factory.New(booking.NewConfig{
			ListingUUID: cmd.ListingUUID,
			People:      *cmd.People,
			User:        cmd.User,
			State:       booking.Created,
			StartDate:   startDate,
			EndDate:     endDate,
			IsPublic:    cmd.IsPublic,
			Listing: booking.Listing{
				Title:        listing.Title,
				Slug:         listing.Slug,
				Description:  listing.Description,
				BusinessName: listing.BusinessName,
				CityName:     listing.CityName,
				DistrictName: listing.DistrictName,
				CountryName:  listing.CountryName,
				Images:       listingImages,
			},
		})
		error := factory.Validate(e)
		if error != nil {
			return nil, error
		}
		res, err := repo.Create(ctx, e)
		if err != nil {
			return nil, err
		}
		events.Created(booking.CreatedEvent{
			BookingUUID: res.UUID,
			ListingUUID: res.ListingUUID,
			People:      &res.People,
			StartDate:   res.StartDate,
			EndDate:     res.EndDate,
		})
		return &BookingCreateRes{
			UUID: res.UUID,
		}, nil
	}
}
