package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/pkg/utils"
)

type BookingListQuery struct {
	*utils.Pagination
	*booking.FilterEntity
	UserUUID string `params:"-" query:"-"`
	UserName string `params:"-" query:"-"`
}

type BookingListRes struct {
	List *list.Result[booking.BookingListDto]
}

type BookingListHandler cqrs.HandlerFunc[BookingListQuery, *BookingListRes]

func NewBookingListHandler(repo booking.Repo) BookingListHandler {
	return func(ctx context.Context, query BookingListQuery) (*BookingListRes, *i18np.Error) {
		query.Default()
		if query.FilterEntity.Type != booking.TypeGuest && query.FilterEntity.Type != booking.TypeOrganizer {
			query.FilterEntity.Type = booking.TypeAny
		}
		query.FilterEntity.UserUUID = query.UserUUID
		query.FilterEntity.UserName = query.UserName
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.List(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]booking.BookingListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToListDto()
		}
		return &BookingListRes{
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
