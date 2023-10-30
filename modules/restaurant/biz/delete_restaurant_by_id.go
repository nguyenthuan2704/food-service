package biz

import (
	"context"
	"food-delivery/modules/restaurant/model"
)

type DeleteRestaurantStorage interface {
	GetRestaurant(ctx context.Context, cond map[string]interface{}) (*model.Restaurant, error)
	DeleteRestaurant(ctx context.Context, cond map[string]interface{}) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStorage
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStorage) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurantById(ctx context.Context, id int) error {
	data, err := biz.store.GetRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.RestaurantDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.store.DeleteRestaurant(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
