package booking

type fieldsType struct {
	UUID       string
	PostUUID   string
	OwnerUUID  string
	People     string
	User       string
	Guests     string
	Days       string
	State      string
	IsPublic   string
	TotalPrice string
	StartDate  string
	EndDate    string
	CreatedAt  string
	UpdatedAt  string
}

type peopleFieldsType struct {
	Adult string
	Kid   string
	Baby  string
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
	UUID:       "_id",
	PostUUID:   "post_uuid",
	OwnerUUID:  "owner_uuid",
	People:     "people",
	User:       "user",
	Guests:     "guests",
	Days:       "days",
	State:      "state",
	IsPublic:   "is_public",
	TotalPrice: "total_price",
	StartDate:  "start_date",
	EndDate:    "end_date",
	CreatedAt:  "createdAt",
	UpdatedAt:  "updatedAt",
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
