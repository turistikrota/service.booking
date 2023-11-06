package invite

type fieldsType struct {
	UUID            string
	BookingUUID     string
	Email           string
	IsUsed          string
	IsDeleted       string
	CreatedAt       string
	UpdatedAt       string
	CreatorUserName string
}

var fields = fieldsType{
	UUID:            "_id",
	BookingUUID:     "booking_uuid",
	Email:           "email",
	IsUsed:          "is_used",
	CreatorUserName: "creator_user_name",
	IsDeleted:       "is_deleted",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}
