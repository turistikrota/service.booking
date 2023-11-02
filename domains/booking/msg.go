package booking

type messages struct {
	Failed        string
	InvalidUUID   string
	InternalError string
}

var I18nMessages = messages{
	Failed:        "booking_failed",
	InvalidUUID:   "invalid_uuid",
	InternalError: "internal_error",
}
