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

func (e *Entity) ToViewDto() BookingViewDto {
	return BookingViewDto{}
}

func (e *Entity) ToAdminViewDto() BookingAdminViewDto {
	return BookingAdminViewDto{}
}
