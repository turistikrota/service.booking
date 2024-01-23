package booking

import "time"

type Entity struct {
	UUID         string             `json:"uuid" bson:"_id,omitempty"`
	ListingUUID  string             `json:"listingUUID" bson:"listing_uuid"`
	BusinessUUID string             `json:"businessUUID" bson:"business_uuid"`
	Listing      Listing            `json:"listing" bson:"listing"`
	People       People             `json:"people" bson:"people"`
	User         User               `json:"user" bson:"user"`
	Guests       []Guest            `json:"guests" bson:"guests"`
	Days         []Day              `json:"days" bson:"days"`
	State        State              `json:"state" bson:"state"`
	Currency     Currency           `json:"currency" bson:"currency"`
	CancelReason *CancelReason      `json:"cancelReason,omitempty" bson:"cancel_reason,omitempty"`
	Errors       []*ValidationError `json:"errors,omitempty" bson:"errors,omitempty"`
	IsPublic     *bool              `json:"isPublic" bson:"is_public"`
	Price        float64            `json:"price" bson:"price"`
	TotalPrice   float64            `json:"totalPrice" bson:"total_price"`
	StartDate    time.Time          `json:"startDate" bson:"start_date"`
	EndDate      time.Time          `json:"endDate" bson:"end_date"`
	CreatedAt    time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updated_at"`
}

type CancelReason struct {
	Content     map[Locale]string `json:"content" bson:"content"`
	CancelledBy CancelOwner       `json:"cancelledBy" bson:"cancelled_by"`
	CancelledAt time.Time         `json:"cancelledAt" bson:"cancelled_at"`
}

type Listing struct {
	Title        string         `json:"title" bson:"title"`
	Slug         string         `json:"slug" bson:"slug"`
	Description  string         `json:"description" bson:"description"`
	BusinessName string         `json:"businessName" bson:"business_name"`
	CityName     string         `json:"cityName" bson:"city_name"`
	DistrictName string         `json:"districtName" bson:"district_name"`
	CountryName  string         `json:"countryName" bson:"country_name"`
	Images       []ListingImage `json:"images" bson:"images"`
}

type ListingImage struct {
	Url   string `json:"url" bson:"url"`
	Order int    `json:"order" bson:"order"`
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

type CancelOwner string

const (
	CancelOwnerAdmin    CancelOwner = "admin"
	CancelOwnerBusiness CancelOwner = "business"
)

type Currency string

const (
	CurrencyTRY Currency = "TRY"
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
)

type Locale string

const (
	LocaleTR Locale = "tr"
	LocaleEN Locale = "en"
)

type State string

const (
	Canceled     State = "canceled"
	NotValid     State = "not_valid"
	Created      State = "created"
	PayExpired   State = "pay_expired"
	PayCancelled State = "pay_cancelled"
	PayPending   State = "pay_pending"
	PayPaid      State = "pay_paid"
	PayRefunded  State = "pay_refunded"
)

func (s State) String() string {
	return string(s)
}
