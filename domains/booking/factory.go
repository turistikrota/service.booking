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
	ListingUUID string
	People      People
	User        User
	State       State
	StartDate   time.Time
	EndDate     time.Time
	IsPublic    *bool
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	return &Entity{
		ListingUUID: cnf.ListingUUID,
		People:      cnf.People,
		User:        cnf.User,
		Guests:      []Guest{},
		Days:        []Day{},
		State:       cnf.State,
		IsPublic:    cnf.IsPublic,
		StartDate:   cnf.StartDate,
		EndDate:     cnf.EndDate,
		CreatedAt:   t,
		UpdatedAt:   t,
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
