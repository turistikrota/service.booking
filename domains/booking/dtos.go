package booking

import "time"

type BookingListDto struct {
	UUID       string    `json:"uuid"`
	People     People    `json:"people"`
	Guests     []Guest   `json:"guests"`
	State      State     `json:"state"`
	IsPublic   *bool     `json:"isPublic"`
	TotalPrice float64   `json:"totalPrice"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

type BookingViewDto struct {
	UUID         string    `json:"uuid"`
	PostUUID     string    `json:"postUUID"`
	BusinessUUID string    `json:"businessUUID"`
	People       People    `json:"people"`
	Guests       []Guest   `json:"guests"`
	Days         []Day     `json:"days"`
	State        State     `json:"state"`
	IsPublic     *bool     `json:"isPublic"`
	TotalPrice   float64   `json:"totalPrice"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BookingAdminViewDto struct {
	UUID         string    `json:"uuid"`
	PostUUID     string    `json:"postUUID"`
	BusinessUUID string    `json:"businessUUID"`
	User         User      `json:"user"`
	People       People    `json:"people"`
	Guests       []Guest   `json:"guests"`
	Days         []Day     `json:"days"`
	State        State     `json:"state"`
	IsPublic     *bool     `json:"isPublic"`
	TotalPrice   float64   `json:"totalPrice"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BookingAdminListDto struct {
	UUID       string    `json:"uuid"`
	User       User      `json:"user"`
	People     People    `json:"people"`
	Guests     []Guest   `json:"guests"`
	State      State     `json:"state"`
	IsPublic   *bool     `json:"isPublic"`
	TotalPrice float64   `json:"totalPrice"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

type BookingBusinessViewDto struct {
	UUID         string    `json:"uuid"`
	PostUUID     string    `json:"postUUID"`
	BusinessUUID string    `json:"businessUUID"`
	People       People    `json:"people"`
	Guests       []Guest   `json:"guests"`
	Days         []Day     `json:"days"`
	State        State     `json:"state"`
	IsPublic     *bool     `json:"isPublic"`
	TotalPrice   float64   `json:"totalPrice"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BookingBusinessListDto struct {
	UUID       string    `json:"uuid"`
	People     People    `json:"people"`
	Guests     []Guest   `json:"guests"`
	State      State     `json:"state"`
	IsPublic   *bool     `json:"isPublic"`
	TotalPrice float64   `json:"totalPrice"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (e *Entity) ToListDto() BookingListDto {
	return BookingListDto{
		UUID:       e.UUID,
		People:     e.People,
		Guests:     e.Guests,
		State:      e.State,
		IsPublic:   e.IsPublic,
		TotalPrice: e.TotalPrice,
		StartDate:  e.StartDate,
		EndDate:    e.EndDate,
		CreatedAt:  e.CreatedAt,
	}
}
func (e *Entity) ToBusinessListDto() BookingBusinessListDto {
	return BookingBusinessListDto{
		UUID:       e.UUID,
		People:     e.People,
		Guests:     e.Guests,
		State:      e.State,
		IsPublic:   e.IsPublic,
		TotalPrice: e.TotalPrice,
		StartDate:  e.StartDate,
		EndDate:    e.EndDate,
		CreatedAt:  e.CreatedAt,
	}
}

func (e *Entity) ToViewDto() BookingViewDto {
	return BookingViewDto{
		UUID:         e.UUID,
		PostUUID:     e.PostUUID,
		BusinessUUID: e.BusinessUUID,
		People:       e.People,
		Guests:       e.Guests,
		State:        e.State,
		IsPublic:     e.IsPublic,
		Days:         e.Days,
		TotalPrice:   e.TotalPrice,
		StartDate:    e.StartDate,
		EndDate:      e.EndDate,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

func (e *Entity) ToAdminViewDto() BookingAdminViewDto {
	return BookingAdminViewDto{
		UUID:         e.UUID,
		PostUUID:     e.PostUUID,
		BusinessUUID: e.BusinessUUID,
		People:       e.People,
		User:         e.User,
		Guests:       e.Guests,
		State:        e.State,
		IsPublic:     e.IsPublic,
		Days:         e.Days,
		TotalPrice:   e.TotalPrice,
		StartDate:    e.StartDate,
		EndDate:      e.EndDate,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

func (e *Entity) ToAdminListDto() BookingAdminListDto {
	return BookingAdminListDto{
		UUID:       e.UUID,
		People:     e.People,
		Guests:     e.Guests,
		State:      e.State,
		IsPublic:   e.IsPublic,
		TotalPrice: e.TotalPrice,
		StartDate:  e.StartDate,
		EndDate:    e.EndDate,
		CreatedAt:  e.CreatedAt,
		User:       e.User,
	}
}
