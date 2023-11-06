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

type BookingListMyAttendeesQuery struct {
	*utils.Pagination
	UserUUID string `params:"-" query:"-"`
	UserName string `params:"-" query:"-"`
}

type BookingListMyAttendeesRes struct {
	List *list.Result[booking.BookingListDto]
}

type BookingListMyAttendeesHandler cqrs.HandlerFunc[BookingListMyAttendeesQuery, *BookingListMyAttendeesRes]

func NewBookingListMyAttendeesHandler(repo booking.Repo, cacheSrv cache.Service) BookingListMyAttendeesHandler {
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
	return func(ctx context.Context, query BookingListMyAttendeesQuery) (*BookingListMyAttendeesRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		cacheHandler := func() (*list.Result[*booking.Entity], *i18np.Error) {
			return repo.ListMyAttendees(ctx, booking.WithUser{
				UUID: query.UserUUID,
				Name: query.UserName,
			}, list.Config{
				Offset: offset,
				Limit:  *query.Limit,
			})
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("booking:by_attendees:name:%s:offset:%v:limit:%v", query.UserName, offset, *query.Limit))
		if err != nil {
			return nil, err
		}
		li := make([]booking.BookingListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToListDto()
		}
		return &BookingListMyAttendeesRes{
			List: &list.Result[booking.BookingListDto]{
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
