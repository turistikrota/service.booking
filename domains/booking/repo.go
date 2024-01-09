package booking

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WithUser struct {
	UUID string
	Name string
}

type Validated struct {
	UUID         string
	ListingUUID  string
	BusinessUUID string
	BusinessName string
	TotalPrice   float64
	Days         []Day
}

type Repo interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	Cancel(ctx context.Context, uuid string) *i18np.Error
	Validated(ctx context.Context, v *Validated) *i18np.Error
	MarkPending(ctx context.Context, uuid string) *i18np.Error
	MarkExpired(ctx context.Context, uuid string) *i18np.Error
	MarkPaid(ctx context.Context, uuid string, totalPrice float64) *i18np.Error
	MarkRefunded(ctx context.Context, uuid string) *i18np.Error
	MarkPayCancelled(ctx context.Context, uuid string) *i18np.Error
	MarkNotValid(ctx context.Context, uuid string) *i18np.Error
	MarkPublic(ctx context.Context, uuid string) *i18np.Error
	MarkPrivate(ctx context.Context, uuid string) *i18np.Error
	GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	View(ctx context.Context, uuid string, userName string) (*Entity, *i18np.Error)
	AddGuest(ctx context.Context, uuid string, guest *Guest) *i18np.Error
	RemoveGuest(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error
	MarkGuestAsPublic(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error
	MarkGuestAsPrivate(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error
	List(ctx context.Context, filter FilterEntity, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	ListByBusiness(ctx context.Context, filter FilterEntity, businessUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	ListByListing(ctx context.Context, filter FilterEntity, listingUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	ListByUser(ctx context.Context, filter FilterEntity, userName string, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	GetDetailWithUser(ctx context.Context, uuid string, userUUID string, userName string) (*Entity, *bool, *i18np.Error)
	CheckAvailability(ctx context.Context, listingUUID string, startDate time.Time, endDate time.Time) (bool, *i18np.Error)
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepo(collection *mongo.Collection, factory Factory) Repo {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func (r *repo) Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error) {
	res, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	entity.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}

func (r *repo) Cancel(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     Canceled,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Validated(ctx context.Context, v *Validated) *i18np.Error {
	id, err := mongo2.TransformId(v.UUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.BusinessUUID: v.BusinessUUID,
			fields.Days:         v.Days,
			fields.Price:        v.TotalPrice,
			fields.State:        PayPending,
			fields.UpdatedAt:    time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkExpired(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     PayExpired,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkPayCancelled(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     PayCancelled,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkPending(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     PayPending,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkPaid(ctx context.Context, uuid string, totalPrice float64) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.TotalPrice: totalPrice,
			fields.State:      PayPaid,
			fields.UpdatedAt:  time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkRefunded(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     PayRefunded,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkNotValid(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     NotValid,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkPublic(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
		fields.IsPublic: bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsPublic:  true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkPrivate(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
		fields.IsPublic: bson.M{
			"$ne": false,
		},
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsPublic:  false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AddGuest(ctx context.Context, uuid string, guest *Guest) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.Guests: bson.M{
				guestFields.UUID:     guest.UUID,
				guestFields.Name:     guest.Name,
				guestFields.IsPublic: guest.IsPublic,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) RemoveGuest(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                  id,
		guestField(guestFields.Name): guest.Name,
		guestField(guestFields.UUID): guest.UUID,
	}
	update := bson.M{
		"$pull": bson.M{
			fields.Guests: bson.M{
				guestFields.UUID: guest.UUID,
				guestFields.Name: guest.Name,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkGuestAsPublic(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                  id,
		guestField(guestFields.Name): guest.Name,
		guestField(guestFields.UUID): guest.UUID,
	}
	update := bson.M{
		"$set": bson.M{
			guestFieldInArray(guestFields.IsPublic): true,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkGuestAsPrivate(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                  id,
		guestField(guestFields.Name): guest.Name,
		guestField(guestFields.UUID): guest.UUID,
	}
	update := bson.M{
		"$set": bson.M{
			guestFieldInArray(guestFields.IsPublic): false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) List(ctx context.Context, filter FilterEntity, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filters := r.filterToBson(filter)
	l, err := r.helper.GetListFilter(ctx, filters, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConf.Offset+listConf.Limit,
		IsPrev:        listConf.Offset > 0,
		FilteredTotal: filtered,
		Total:         filtered,
		Page:          listConf.Offset/listConf.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) ListByBusiness(ctx context.Context, filter FilterEntity, businessUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter.Type = ""
	filters := r.filterToBson(filter, bson.M{
		fields.BusinessUUID: businessUUID,
	})
	l, err := r.helper.GetListFilter(ctx, filters, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConf.Offset+listConf.Limit,
		IsPrev:        listConf.Offset > 0,
		FilteredTotal: filtered,
		Total:         filtered,
		Page:          listConf.Offset/listConf.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) ListByListing(ctx context.Context, filter FilterEntity, listingUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter.Type = ""
	filters := r.filterToBson(filter, bson.M{
		fields.ListingUUID: listingUUID,
	})
	l, err := r.helper.GetListFilter(ctx, filters, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConf.Offset+listConf.Limit,
		IsPrev:        listConf.Offset > 0,
		FilteredTotal: filtered,
		Total:         filtered,
		Page:          listConf.Offset/listConf.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) ListByUser(ctx context.Context, filter FilterEntity, userName string, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter.Type = ""
	filters := r.filterToBson(filter, bson.M{
		"$or": []bson.M{
			{
				userField(userFields.Name): userName,
			},
			{
				guestField(guestFields.Name): userName,
			},
		},
	})
	l, err := r.helper.GetListFilter(ctx, filters, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConf.Offset+listConf.Limit,
		IsPrev:        listConf.Offset > 0,
		FilteredTotal: filtered,
		Total:         filtered,
		Page:          listConf.Offset/listConf.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) GetDetailWithUser(ctx context.Context, uuid string, userUUID string, userName string) (*Entity, *bool, *i18np.Error) {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return nil, nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                id,
		userField(userFields.UUID): userUUID,
		userField(userFields.Name): userName,
	}
	res, exists, _err := r.helper.GetFilter(ctx, filter)
	if _err != nil {
		return nil, nil, r.factory.Errors.InternalError()
	}
	return *res, &exists, nil
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	res, _, _err := r.helper.GetFilter(ctx, filter)
	if _err != nil {
		return nil, r.factory.Errors.InternalError()
	}
	return *res, nil
}

func (r *repo) View(ctx context.Context, uuid string, userName string) (*Entity, *i18np.Error) {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		"$or": []bson.M{
			{
				fields.UUID: id,
				fields.IsPublic: bson.M{
					"$eq": true,
				},
			},
			{
				fields.UUID:                id,
				userField(userFields.Name): userName,
			},
			{
				fields.UUID:                  id,
				guestField(guestFields.Name): userName,
			},
		},
	}
	res, exists, _err := r.helper.GetFilter(ctx, filter)
	if _err != nil {
		return nil, r.factory.Errors.InternalError()
	}
	if !exists {
		return nil, r.factory.Errors.NotFound()
	}
	return *res, nil
}

func (r *repo) CheckAvailability(ctx context.Context, listingUUID string, startDate time.Time, endDate time.Time) (bool, *i18np.Error) {
	filter := bson.M{
		fields.ListingUUID: listingUUID,
		fields.State: bson.M{
			"$in": []State{
				Created,
				PayPending,
				PayPaid,
			},
		},
		fields.StartDate: bson.M{
			"$lte": endDate,
		},
		"$or": []bson.M{
			{
				fields.EndDate: bson.M{
					"$gt": startDate,
				},
			},
			{
				fields.StartDate: startDate,
				fields.EndDate: bson.M{
					"$gte": endDate,
				},
			},
		},
	}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, r.factory.Errors.InternalError()
	}
	return count == 0, nil
}

func (r *repo) listOptions(listConfig list.Config) *options.FindOptions {
	opts := options.Find()
	return opts.SetLimit(listConfig.Limit).SetSkip(listConfig.Offset)
}
