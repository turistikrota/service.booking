package config

import "github.com/turistikrota/service.shared/base_roles"

type bookingRoles struct {
	List string
	View string
}

type roles struct {
	base_roles.Roles
	Booking bookingRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Booking: bookingRoles{
		List: "booking.list",
		View: "booking.view",
	},
}
