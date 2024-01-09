package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/pkg/utils"
)

type BookingListByListingQuery struct {
	*utils.Pagination
	*booking.FilterEntity
	ListingUUID string `params:"uuid" query:"-" validate:"required,object_id"`
	IsPublic    bool   `params:"-" query:"-"`
}

type BookingListByListingRes struct {
	List *list.Result[booking.BookingBusinessListDto]
}

type BookingListByListingHandler cqrs.HandlerFunc[BookingListByListingQuery, *BookingListByListingRes]

func NewBookingListByListingHandler(repo booking.Repo, cacheSrv cache.Service) BookingListByListingHandler {
	cache := cache.New[*list.Result[*booking.Entity]](cacheSrv)

	createCacheEntity := func() *list.Result[*booking.Entity] {
		return &list.Result[*booking.Entity]{
			List:          []*booking.Entity{},
			Total:         0,
			FilteredTotal: 0,
			Page:          0,
			IsNext:        false,
			IsPrev:        false,
		}
	}
	return func(ctx context.Context, query BookingListByListingQuery) (*BookingListByListingRes, *i18np.Error) {
		query.Default()
		query.FilterEntity.ForPrivate()
		if query.IsPublic {
			query.FilterEntity.PublicView()
		}
		offset := (*query.Page - 1) * *query.Limit
		cacheHandler := func() (*list.Result[*booking.Entity], *i18np.Error) {
			return repo.ListByListing(ctx, *query.FilterEntity, query.ListingUUID, list.Config{
				Offset: offset,
				Limit:  *query.Limit,
			})
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("booking:by_listing:uuid:%s:offset:%v:limit:%v", query.ListingUUID, offset, *query.Limit))
		if err != nil {
			return nil, err
		}
		li := make([]booking.BookingBusinessListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToBusinessListDto()
		}
		return &BookingListByListingRes{
			List: &list.Result[booking.BookingBusinessListDto]{
				List:          li,
				Total:         res.Total,
				FilteredTotal: res.FilteredTotal,
				Page:          res.Page,
				IsNext:        res.IsNext,
				IsPrev:        res.IsPrev,
			},
		}, nil
	}
}
