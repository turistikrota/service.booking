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
	PostUUID  string
	People    People
	User      User
	State     State
	StartDate time.Time
	EndDate   time.Time
	IsPublic  *bool
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	return &Entity{
		PostUUID:  cnf.PostUUID,
		People:    cnf.People,
		User:      cnf.User,
		Guests:    []Guest{},
		Days:      []Day{},
		State:     cnf.State,
		IsPublic:  cnf.IsPublic,
		StartDate: cnf.StartDate,
		EndDate:   cnf.EndDate,
		CreatedAt: t,
		UpdatedAt: t,
	}
}
