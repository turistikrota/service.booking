package booking

type fieldsType struct {
	UUID         string
	ListingUUID  string
	BusinessUUID string
	Listing      string
	People       string
	User         string
	Guests       string
	CancelReason string
	Errors       string
	Days         string
	State        string
	IsPublic     string
	Price        string
	Currency     string
	TotalPrice   string
	StartDate    string
	EndDate      string
	CreatedAt    string
	UpdatedAt    string
}

type peopleFieldsType struct {
	Adult string
	Kid   string
	Baby  string
}

type listingFieldsType struct {
	Title        string
	Slug         string
	Description  string
	BusinessName string
	CityName     string
	DistrictName string
	CountryName  string
	Images       string
}

type userFieldsType struct {
	UUID string
	Name string
}

type guestFieldsType struct {
	UUID     string
	Name     string
	IsPublic string
}

type dayFieldsType struct {
	Date  string
	Price string
}

var fields = fieldsType{
	UUID:         "_id",
	ListingUUID:  "listing_uuid",
	BusinessUUID: "business_uuid",
	Listing:      "listing",
	People:       "people",
	User:         "user",
	CancelReason: "cancel_reason",
	Guests:       "guests",
	Errors:       "errors",
	Days:         "days",
	State:        "state",
	Currency:     "currency",
	IsPublic:     "is_public",
	Price:        "price",
	TotalPrice:   "total_price",
	StartDate:    "start_date",
	EndDate:      "end_date",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var listingFields = listingFieldsType{
	Title:        "title",
	Slug:         "slug",
	Description:  "description",
	BusinessName: "business_name",
	CityName:     "city_name",
	DistrictName: "district_name",
	CountryName:  "country_name",
	Images:       "images",
}

var peopleFields = peopleFieldsType{
	Adult: "adult",
	Kid:   "kid",
	Baby:  "baby",
}

var userFields = userFieldsType{
	UUID: "uuid",
	Name: "name",
}

var guestFields = guestFieldsType{
	UUID:     "uuid",
	Name:     "name",
	IsPublic: "is_public",
}

func peopleField(key string) string {
	return fields.People + "." + key
}

func userField(key string) string {
	return fields.User + "." + key
}

func guestField(key string) string {
	return fields.Guests + "." + key
}

func guestFieldInArray(key string) string {
	return fields.Guests + ".$." + key
}

func dayField(key string) string {
	return fields.Days + "." + key
}

func listingField(key string) string {
	return fields.Listing + "." + key
}
