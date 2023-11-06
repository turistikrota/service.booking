package booking

type BookingListDto struct{}

type BookingViewDto struct{}

type BookingAdminViewDto struct{}

type BookingAdminListDto struct{}

func (e *Entity) ToListDto() BookingListDto {
	return BookingListDto{}
}
