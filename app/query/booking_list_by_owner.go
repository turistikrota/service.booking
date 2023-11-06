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

type BookingListByOwnerQuery struct {
	*utils.Pagination
	OwnerUUID string `params:"-" query:"-"`
}

type BookingListByOwnerRes struct {
	List *list.Result[booking.BookingOwnerListDto]
}

type BookingListByOwnerHandler cqrs.HandlerFunc[BookingListByOwnerQuery, *BookingListByOwnerRes]

func NewBookingListByOwnerHandler(repo booking.Repo, cacheSrv cache.Service) BookingListByOwnerHandler {
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
	return func(ctx context.Context, query BookingListByOwnerQuery) (*BookingListByOwnerRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		cacheHandler := func() (*list.Result[*booking.Entity], *i18np.Error) {
			return repo.ListByOwner(ctx, query.OwnerUUID, list.Config{
				Offset: offset,
				Limit:  *query.Limit,
			})
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("booking:by_owner:uuid:%s:offset:%v:limit:%v", query.OwnerUUID, offset, *query.Limit))
		if err != nil {
			return nil, err
		}
		li := make([]booking.BookingOwnerListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToOwnerListDto()
		}
		return &BookingListByOwnerRes{
			List: &list.Result[booking.BookingOwnerListDto]{
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
