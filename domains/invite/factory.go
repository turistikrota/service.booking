package invite

import "time"

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{Errors: newInviteErrors()}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) New(email string, bookingUUID string, userName string) *Entity {
	t := time.Now()
	return &Entity{
		Email:           email,
		BookingUUID:     bookingUUID,
		CreatorUserName: userName,
		IsUsed:          false,
		IsDeleted:       false,
		CreatedAt:       &t,
		UpdatedAt:       &t,
	}
}
