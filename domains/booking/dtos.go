package booking

import "time"

type BookingListDto struct {
	UUID       string    `json:"uuid"`
	People     People    `json:"people"`
	Listing    Listing   `json:"listing"`
	Guests     []Guest   `json:"guests"`
	State      State     `json:"state"`
	IsPublic   *bool     `json:"isPublic"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"totalPrice,omitempty"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

type BookingViewDto struct {
	UUID         string    `json:"uuid"`
	ListingUUID  string    `json:"listingUUID"`
	BusinessUUID string    `json:"businessUUID"`
	User         User      `json:"user"`
	Listing      Listing   `json:"listing"`
	People       People    `json:"people"`
	Guests       []Guest   `json:"guests"`
	Days         []Day     `json:"days"`
	State        State     `json:"state"`
	IsPublic     *bool     `json:"isPublic"`
	Price        float64   `json:"price"`
	TotalPrice   float64   `json:"totalPrice,omitempty"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BookingAdminViewDto struct {
	UUID         string    `json:"uuid"`
	ListingUUID  string    `json:"listingUUID"`
	BusinessUUID string    `json:"businessUUID"`
	User         User      `json:"user"`
	Listing      Listing   `json:"listing"`
	People       People    `json:"people"`
	Guests       []Guest   `json:"guests"`
	Days         []Day     `json:"days"`
	State        State     `json:"state"`
	IsPublic     *bool     `json:"isPublic"`
	Price        float64   `json:"price"`
	TotalPrice   float64   `json:"totalPrice,omitempty"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BookingAdminListDto struct {
	UUID       string    `json:"uuid"`
	User       User      `json:"user"`
	People     People    `json:"people"`
	Listing    Listing   `json:"listing"`
	Guests     []Guest   `json:"guests"`
	State      State     `json:"state"`
	IsPublic   *bool     `json:"isPublic"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"totalPrice,omitempty"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

type BookingBusinessViewDto struct {
	UUID         string    `json:"uuid"`
	ListingUUID  string    `json:"listingUUID"`
	BusinessUUID string    `json:"businessUUID"`
	Listing      Listing   `json:"listing"`
	People       People    `json:"people"`
	Guests       []Guest   `json:"guests"`
	Days         []Day     `json:"days"`
	State        State     `json:"state"`
	IsPublic     *bool     `json:"isPublic"`
	Price        float64   `json:"price"`
	TotalPrice   float64   `json:"totalPrice,omitempty"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BookingBusinessListDto struct {
	UUID       string    `json:"uuid"`
	People     People    `json:"people"`
	Listing    Listing   `json:"listing"`
	Guests     []Guest   `json:"guests"`
	State      State     `json:"state"`
	IsPublic   *bool     `json:"isPublic"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"totalPrice,omitempty"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (e *Entity) ToListDto() BookingListDto {
	return BookingListDto{
		UUID:       e.UUID,
		People:     e.People,
		Listing:    e.Listing,
		Guests:     e.Guests,
		State:      e.State,
		IsPublic:   e.IsPublic,
		Price:      e.Price,
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
		Listing:    e.Listing,
		Guests:     e.Guests,
		State:      e.State,
		IsPublic:   e.IsPublic,
		Price:      e.Price,
		TotalPrice: e.TotalPrice,
		StartDate:  e.StartDate,
		EndDate:    e.EndDate,
		CreatedAt:  e.CreatedAt,
	}
}

func (e *Entity) ToViewDto(userId string, userName string) BookingViewDto {
	guests := make([]Guest, 0)
	if userId == e.User.UUID && userName == e.User.Name {
		guests = e.Guests
	} else {
		for _, guest := range e.Guests {
			if guest.IsPublic {
				guests = append(guests, guest)
			}
		}
	}
	return BookingViewDto{
		UUID:         e.UUID,
		ListingUUID:  e.ListingUUID,
		BusinessUUID: e.BusinessUUID,
		Listing:      e.Listing,
		User:         e.User,
		People:       e.People,
		Guests:       guests,
		State:        e.State,
		IsPublic:     e.IsPublic,
		Days:         e.Days,
		Price:        e.Price,
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
		ListingUUID:  e.ListingUUID,
		BusinessUUID: e.BusinessUUID,
		Listing:      e.Listing,
		People:       e.People,
		User:         e.User,
		Guests:       e.Guests,
		State:        e.State,
		IsPublic:     e.IsPublic,
		Days:         e.Days,
		Price:        e.Price,
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
		Listing:    e.Listing,
		State:      e.State,
		IsPublic:   e.IsPublic,
		Price:      e.Price,
		TotalPrice: e.TotalPrice,
		StartDate:  e.StartDate,
		EndDate:    e.EndDate,
		CreatedAt:  e.CreatedAt,
		User:       e.User,
	}
}
