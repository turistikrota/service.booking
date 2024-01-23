package booking

import (
	"time"

	"github.com/cilloparch/cillop/i18np"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: NewErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	ListingUUID  string
	BusinessUUID string
	People       People
	User         User
	State        State
	Listing      Listing
	StartDate    time.Time
	EndDate      time.Time
	IsPublic     *bool
}

type NewCancelConfig struct {
	TrContent string
	EnContent string
	IsAdmin   bool
}

func (f Factory) NewCancelReason(cnf NewCancelConfig) *CancelReason {
	cancelledBy := CancelOwnerBusiness
	if cnf.IsAdmin {
		cancelledBy = CancelOwnerAdmin
	}
	return &CancelReason{
		Content: map[Locale]string{
			LocaleTR: cnf.TrContent,
			LocaleEN: cnf.EnContent,
		},
		CancelledBy: cancelledBy,
		CancelledAt: time.Now(),
	}
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	return &Entity{
		ListingUUID:  cnf.ListingUUID,
		BusinessUUID: cnf.BusinessUUID,
		People:       cnf.People,
		User:         cnf.User,
		Listing:      cnf.Listing,
		Guests:       []Guest{},
		Days:         []Day{},
		State:        cnf.State,
		IsPublic:     cnf.IsPublic,
		StartDate:    cnf.StartDate,
		EndDate:      cnf.EndDate,
		CreatedAt:    t,
		UpdatedAt:    t,
	}
}

type validator func(e *Entity) *i18np.Error

func (f Factory) Validate(e *Entity) *i18np.Error {
	validators := []validator{
		f.validateTime,
	}
	for _, v := range validators {
		if err := v(e); err != nil {
			return err
		}
	}
	return nil
}

func (f Factory) validateTime(e *Entity) *i18np.Error {
	return f.ValidateDateTime(e.StartDate, e.EndDate)
}

func (f Factory) ParseDatesWithoutHours(startDate time.Time, endDate time.Time) (time.Time, time.Time) {
	return f.parseDateWithoutHours(startDate), f.parseDateWithoutHours(endDate)
}

func (f Factory) parseDateWithoutHours(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

func (f Factory) ValidateDateTime(startDate time.Time, endDate time.Time) *i18np.Error {
	if startDate.After(endDate) {
		return f.Errors.StartDateAfterEndDate()
	}
	if startDate.Before(time.Now()) {
		return f.Errors.StartDateBeforeNow()
	}
	return nil
}

func (f Factory) IsCancelable(e *Entity) bool {
	disallowStatus := []State{
		Canceled,
		PayRefunded,
	}
	for _, s := range disallowStatus {
		if e.State == s {
			return false
		}
	}
	now := time.Now()
	return e.StartDate.After(now) && e.EndDate.After(now)
}

func (f Factory) IsCancelableAsBusiness(e *Entity) bool {
	disallowStatus := []State{
		Canceled,
		PayRefunded,
	}
	for _, s := range disallowStatus {
		if e.State == s {
			return false
		}
	}
	now := time.Now()
	return now.Before(e.StartDate) && e.CreatedAt.Add(72*time.Hour).After(now)
}
