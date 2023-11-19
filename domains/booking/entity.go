package booking

import "time"

type Entity struct {
	UUID         string    `json:"uuid" bson:"_id,omitempty"`
	PostUUID     string    `json:"postUUID" bson:"post_uuid"`
	BusinessUUID string    `json:"businessUUID" bson:"business_uuid"`
	People       People    `json:"people" bson:"people"`
	User         User      `json:"user" bson:"user"`
	Guests       []Guest   `json:"guests" bson:"guests"`
	Days         []Day     `json:"days" bson:"days"`
	State        State     `json:"state" bson:"state"`
	IsPublic     *bool     `json:"isPublic" bson:"is_public"`
	TotalPrice   float64   `json:"totalPrice" bson:"total_price"`
	StartDate    time.Time `json:"startDate" bson:"start_date"`
	EndDate      time.Time `json:"endDate" bson:"end_date"`
	CreatedAt    time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updated_at"`
}

type User struct {
	UUID string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
}

type Guest struct {
	UUID     string `json:"uuid" bson:"uuid"`
	Name     string `json:"name" bson:"name"`
	IsPublic bool   `json:"isPublic" bson:"is_public"`
}

type Day struct {
	Date  time.Time `json:"date" bson:"date"`
	Price float64   `json:"price" bson:"price"`
}

type People struct {
	Adult int `json:"adult" bson:"adult" validate:"required,gt=0"`
	Kid   int `json:"kid" bson:"kid" validate:"gte=0"`
	Baby  int `json:"baby" bson:"baby" validate:"gte=0"`
}

type State string

const (
	Canceled State = "canceled"
	NotValid State = "not_valid"
	Created  State = "created"
	Expired  State = "expired"
	Pending  State = "pending"
	Paid     State = "paid"
	Refunded State = "refunded"
)

func (s State) String() string {
	return string(s)
}
