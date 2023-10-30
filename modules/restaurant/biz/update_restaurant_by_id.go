package biz

import (
	"context"
	"food-delivery/modules/restaurant/model"
)

type UpdateRestaurantStorage interface {
	GetRestaurant(ctx context.Context, cond map[string]interface{}) (*model.Restaurant, error)
	UpdateRestaurant(ctx context.Context, cond map[string]interface{}, dataUpdate *model.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStorage
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStorage) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurantById(ctx context.Context, id int, dataUpdate *model.RestaurantUpdate) error {
	data, err := biz.store.GetRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.RestaurantDeleted {
		/*return model.ErrItemDeleted*/
		return biz.store.UpdateRestaurant(ctx, map[string]interface{}{"id": id}, dataUpdate)
	}

	if err := biz.store.UpdateRestaurant(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
