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

type Repo interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	Cancel(ctx context.Context, uuid string) *i18np.Error
	MarkPending(ctx context.Context, uuid string) *i18np.Error
	MarkExpired(ctx context.Context, uuid string) *i18np.Error
	MarkPaid(ctx context.Context, uuid string) *i18np.Error
	MarkRefunded(ctx context.Context, uuid string) *i18np.Error
	MarkUsed(ctx context.Context, uuid string) *i18np.Error
	MarkPublic(ctx context.Context, uuid string) *i18np.Error
	MarkPrivate(ctx context.Context, uuid string) *i18np.Error
	AddGuest(ctx context.Context, uuid string, guest *Guest) *i18np.Error
	RemoveGuest(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error
	MarkGuestAsPublic(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error
	MarkGuestAsPrivate(ctx context.Context, uuid string, guest WithUser, user WithUser) *i18np.Error
	ListMyOrganized(ctx context.Context, user WithUser, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	ListMyAttendees(ctx context.Context, user WithUser, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	ListByOwner(ctx context.Context, ownerUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	ListByPost(ctx context.Context, postUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
	GetDetailWithUser(ctx context.Context, uuid string, userUUID string) (*Entity, *i18np.Error)
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
			fields.State:     Expired,
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
			fields.State:     Pending,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkPaid(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     Paid,
			fields.UpdatedAt: time.Now(),
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
			fields.State:     Refunded,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkUsed(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.State:     Used,
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
				guestField(guestFields.Name): guest.Name,
				guestField(guestFields.UUID): guest.UUID,
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

func (r *repo) ListMyOrganized(ctx context.Context, user WithUser, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter := bson.M{
		userField(userFields.UUID): user.UUID,
		userField(userFields.Name): user.Name,
	}
	l, err := r.helper.GetListFilter(ctx, filter, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filter)
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

func (r *repo) ListMyAttendees(ctx context.Context, user WithUser, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter := bson.M{
		guestField(guestFields.UUID): user.UUID,
		guestField(guestFields.Name): user.Name,
	}
	l, err := r.helper.GetListFilter(ctx, filter, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filter)
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

func (r *repo) ListByOwner(ctx context.Context, ownerUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter := bson.M{
		fields.OwnerUUID: ownerUUID,
	}
	l, err := r.helper.GetListFilter(ctx, filter, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filter)
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

func (r *repo) ListByPost(ctx context.Context, postUUID string, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter := bson.M{
		fields.PostUUID: postUUID,
	}
	l, err := r.helper.GetListFilter(ctx, filter, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filter)
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

func (r *repo) GetDetailWithUser(ctx context.Context, uuid string, userUUID string) (*Entity, *i18np.Error) {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:                id,
		userField(userFields.UUID): userUUID,
	}
	res, _, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.InternalError()
	}
	return *res, nil
}

func (r *repo) listOptions(listConfig list.Config) *options.FindOptions {
	opts := options.Find()
	opts.SetLimit(listConfig.Limit).SetSkip(listConfig.Offset)
	return opts
}
