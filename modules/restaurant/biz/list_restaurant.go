package biz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/model"
)

type ListRestaurantStorage interface {
	ListRestaurant(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStorage
}

func NewListRestaurantBiz(store ListRestaurantStorage) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.Restaurant, error) {
	data, err := biz.store.ListRestaurant(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil

}
