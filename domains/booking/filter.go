package booking

import (
	"go.mongodb.org/mongo-driver/bson"
)

type FilterEntity struct {
	Locale   string `json:"-" query:"-"`
	UserName string `json:"-" query:"-"`
	UserUUID string `json:"-" query:"-"`
	Query    string `query:"q,omitempty" validate:"omitempty,max=100"`
	State    string `query:"state,omitempty" validate:"omitempty,oneof=canceled not_valid created pay_expired pay_cancelled pay_pending pay_paid pay_refunded"`
	Type     string `query:"type,omitempty" validate:"omitempty,oneof=guest organizer any"`
	IsPublic *bool  `query:"isPublic,omitempty" validate:"omitempty"`
}

const (
	TypeGuest     = "guest"
	TypeOrganizer = "organizer"
	TypeAny       = "any"
)

func (e *FilterEntity) ForPrivate() *FilterEntity {
	e.Type = ""
	return e
}

func (e FilterEntity) PublicView() *FilterEntity {
	isPublic := true
	e.IsPublic = &isPublic
	return &e
}

func (r *repo) filterToBson(filter FilterEntity, defaultFilters ...bson.M) bson.M {
	list := make([]bson.M, 0)
	if len(defaultFilters) > 0 {
		list = append(list, defaultFilters...)
	}
	list = r.filterByQuery(list, filter)
	list = r.filterByType(list, filter)
	list = r.filterByState(list, filter)
	list = r.filterByIsPublic(list, filter)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}

func (r *repo) filterByIsPublic(list []bson.M, filter FilterEntity) []bson.M {
	if filter.IsPublic != nil {
		list = append(list, bson.M{
			guestField(guestFields.IsPublic): *filter.IsPublic,
		})
	}
	return list
}

func (r *repo) filterByType(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Type != "" {
		if filter.Type == TypeGuest {
			list = append(list, bson.M{
				guestField(userFields.UUID): filter.UserUUID,
				guestField(userFields.Name): filter.UserName,
			})
		}
		if filter.Type == TypeOrganizer {
			list = append(list, bson.M{
				userField(guestFields.UUID): filter.UserUUID,
				userField(guestFields.Name): filter.UserName,
			})
		}
		if filter.Type == TypeAny {
			list = append(list, bson.M{
				"$or": []bson.M{
					{
						guestField(userFields.UUID): filter.UserUUID,
						guestField(userFields.Name): filter.UserName,
					},
					{
						userField(guestFields.UUID): filter.UserUUID,
						userField(guestFields.Name): filter.UserName,
					},
				},
			})
		}
	}
	return list
}

func (r *repo) filterByQuery(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					listingField(listingFields.Title): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					listingField(listingFields.Description): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					listingField(listingFields.BusinessName): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					listingField(listingFields.CityName): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					listingField(listingFields.DistrictName): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					listingField(listingFields.CountryName): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}

func (r *repo) filterByState(list []bson.M, filter FilterEntity) []bson.M {
	if filter.State != "" {
		list = append(list, bson.M{
			fields.State: filter.State,
		})
	}
	return list
}
