package booking

import "time"

type Entity struct {
	UUID       string    `json:"uuid"`
	PostUUID   string    `json:"postUUID"`
	OwnerUUID  string    `json:"ownerUUID"`
	People     People    `json:"people"`
	User       User      `json:"user"`
	Guests     []Guest   `json:"guests"`
	Days       []Day     `json:"days"`
	State      State     `json:"state"`
	IsPublic   *bool     `json:"isPublic"`
	TotalPrice float64   `json:"totalPrice"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type User struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Guest struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	IsPublic bool   `json:"isPublic"`
}

type Day struct {
	Date  time.Time `json:"date"`
	Price float64   `json:"price"`
}

type People struct {
	Adult int `json:"adult" validate:"required,gt=0"`
	Kid   int `json:"kid" validate:"gte=0"`
	Baby  int `json:"baby" validate:"gte=0"`
}

type State string

const (
	Canceled State = "canceled"
	Created  State = "created"
	Expired  State = "expired"
	Pending  State = "pending"
	Paid     State = "paid"
	Refunded State = "refunded"
	Used     State = "used"
)

func (s State) String() string {
	return string(s)
}
