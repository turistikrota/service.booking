package booking

type BookingListDto struct{}

type BookingViewDto struct{}

type BookingAdminViewDto struct{}

type BookingAdminListDto struct{}

type BookingOwnerViewDto struct{}

type BookingOwnerListDto struct{}

func (e *Entity) ToListDto() BookingListDto {
	return BookingListDto{}
}
func (e *Entity) ToOwnerListDto() BookingOwnerListDto {
	return BookingOwnerListDto{}
}
