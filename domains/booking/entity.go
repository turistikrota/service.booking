package booking

import "time"

type Entity struct {
	UUID      string    `json:"uuid"`
	PostUUID  string    `json:"postUUID"`
	People    People    `json:"people"`
	User      User      `json:"user"`
	Guests    []Guest   `json:"guests"`
	Days      []Day     `json:"days"`
	State     State     `json:"state"`
	IsPublic  *bool     `json:"isPublic"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
	Adult int `json:"adult"`
	Kid   int `json:"kid"`
	Baby  int `json:"baby"`
}

type State string

const (
	Canceled State = "canceled"
	Expired  State = "expired"
	Pending  State = "pending"
	Paid     State = "paid"
	Refunded State = "refunded"
	Used     State = "used"
)

func (s State) String() string {
	return string(s)
}
