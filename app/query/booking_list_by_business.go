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

type BookingListByBusinessQuery struct {
	*utils.Pagination
	*booking.FilterEntity
	BusinessUUID string `params:"-" query:"-"`
	IsPublic     bool   `params:"-" query:"-"`
}

type BookingListByBusinessRes struct {
	List *list.Result[booking.BookingBusinessListDto]
}

type BookingListByBusinessHandler cqrs.HandlerFunc[BookingListByBusinessQuery, *BookingListByBusinessRes]

func NewBookingListByBusinessHandler(repo booking.Repo, cacheSrv cache.Service) BookingListByBusinessHandler {
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
	return func(ctx context.Context, query BookingListByBusinessQuery) (*BookingListByBusinessRes, *i18np.Error) {
		query.Default()
		query.FilterEntity.ForPrivate()
		if query.IsPublic {
			query.FilterEntity.PublicView()
		}
		offset := (*query.Page - 1) * *query.Limit
		cacheHandler := func() (*list.Result[*booking.Entity], *i18np.Error) {
			return repo.ListByBusiness(ctx, *query.FilterEntity, query.BusinessUUID, list.Config{
				Offset: offset,
				Limit:  *query.Limit,
			})
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("booking:by_business:uuid:%s:offset:%v:limit:%v", query.BusinessUUID, offset, *query.Limit))
		if err != nil {
			return nil, err
		}
		li := make([]booking.BookingBusinessListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToBusinessListDto()
		}
		return &BookingListByBusinessRes{
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
