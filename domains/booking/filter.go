package booking

import "go.mongodb.org/mongo-driver/bson"

type FilterEntity struct {
	Locale string `json:"-" query:"-"`
	Query  string `query:"q,omitempty" validate:"omitempty,max=100"`
	State  string `query:"state,omitempty" validate:"omitempty,oneof=canceled not_valid created pay_expired pay_cancelled pay_pending pay_paid pay_refunded"`
}

func (r *repo) filterToBson(filter FilterEntity) bson.M {
	list := make([]bson.M, 0)
	list = r.filterByQuery(list, filter)
	list = r.filterByState(list, filter)
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
