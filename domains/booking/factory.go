package booking

import "time"

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
	PostUUID string
	People   People
	User     User
	Days     []Day
	State    State
	IsPublic *bool
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	return &Entity{
		PostUUID:  cnf.PostUUID,
		People:    cnf.People,
		User:      cnf.User,
		Guests:    []Guest{},
		Days:      cnf.Days,
		State:     cnf.State,
		IsPublic:  cnf.IsPublic,
		CreatedAt: t,
		UpdatedAt: t,
	}
}
