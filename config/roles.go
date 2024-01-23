package config

import "github.com/turistikrota/service.shared/base_roles"

type bookingRoles struct {
	List   string
	View   string
	Cancel string
	Super  string
}

type businessRoles struct {
	Super string
}

type roles struct {
	base_roles.Roles
	Booking  bookingRoles
	Business businessRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Booking: bookingRoles{
		List:   "booking.list",
		View:   "booking.view",
		Cancel: "booking.cancel",
		Super:  "booking.super",
	},
	Business: businessRoles{
		Super: "business.super",
	},
}
