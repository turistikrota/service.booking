package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.booking/domains/booking"
	"github.com/turistikrota/service.booking/pkg/utils"
)

type BookingAdminListQuery struct {
	*utils.Pagination
	*booking.FilterEntity
}

type BookingAdminListRes struct {
	List *list.Result[booking.BookingAdminListDto]
}

type BookingAdminListHandler cqrs.HandlerFunc[BookingAdminListQuery, *BookingAdminListRes]

func NewBookingAdminListHandler(repo booking.Repo) BookingAdminListHandler {
	return func(ctx context.Context, query BookingAdminListQuery) (*BookingAdminListRes, *i18np.Error) {
		query.Default()
		query.FilterEntity.ForPrivate()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.List(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]booking.BookingAdminListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToAdminListDto()
		}
		return &BookingAdminListRes{
			List: &list.Result[booking.BookingAdminListDto]{
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
