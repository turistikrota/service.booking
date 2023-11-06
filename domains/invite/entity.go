package invite

import "time"

type Entity struct {
	UUID            string     `json:"uuid"`
	BookingUUID     string     `json:"bookingUUID"`
	CreatorUserName string     `json:"creatorUserName"`
	Email           string     `json:"email"`
	IsUsed          bool       `json:"isUsed"`
	IsDeleted       bool       `json:"isDeleted"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
