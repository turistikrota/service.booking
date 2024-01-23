package booking

type messages struct {
	Failed                   string
	InvalidUUID              string
	InternalError            string
	ListingNotActive         string
	NotAvailable             string
	StartDateAfterEndDate    string
	StartDateBeforeNow       string
	OnlyAdminCanDoThisAction string
	NotCancelable            string
	NotFound                 string
}

var I18nMessages = messages{
	Failed:                   "booking_failed",
	InvalidUUID:              "invalid_uuid",
	InternalError:            "internal_error",
	NotAvailable:             "not_available",
	ListingNotActive:         "listing_not_active",
	StartDateAfterEndDate:    "start_date_after_end_date",
	StartDateBeforeNow:       "start_date_before_now",
	OnlyAdminCanDoThisAction: "only_admin_can_do_this_action",
	NotCancelable:            "not_cancelable",
	NotFound:                 "not_found",
}
