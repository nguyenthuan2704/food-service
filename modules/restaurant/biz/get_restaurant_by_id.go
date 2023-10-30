package biz

import (
	"context"
	"food-delivery/modules/restaurant/model"
)

type GetRestaurantStorage interface {
	GetRestaurant(ctx context.Context, cond map[string]interface{}) (*model.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStorage
}

func NewGetRestaurantBiz(store GetRestaurantStorage) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurantById(ctx context.Context, id int) (*model.Restaurant, error) {
	data, err := biz.store.GetRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
