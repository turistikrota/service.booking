package invite

import "time"

type Entity struct {
	UUID            string     `json:"uuid" bson:"_id,omitempty"`
	BookingUUID     string     `json:"bookingUUID" bson:"booking_uuid"`
	CreatorUserName string     `json:"creatorUserName" bson:"creator_user_name"`
	Email           string     `json:"email" bson:"email"`
	IsUsed          bool       `json:"isUsed" bson:"is_used"`
	IsDeleted       bool       `json:"isDeleted" bson:"is_deleted"`
	CreatedAt       *time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt       *time.Time `json:"updatedAt" bson:"updated_at"`
}
