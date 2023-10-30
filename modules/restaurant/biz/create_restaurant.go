package biz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/model"
	"strings"
)

type CreateRestaurantStorage interface {
	CreateRestaurant(ctx context.Context, data *model.RestaurantCreation) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStorage
}

func NewCreateRestaurantBiz(store CreateRestaurantStorage) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateNewRestaurant(ctx context.Context, data *model.RestaurantCreation) error {
	if strings.TrimSpace(data.Name) == "" {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.CreateRestaurant(ctx, data); err != nil {
		/*return err*/
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
