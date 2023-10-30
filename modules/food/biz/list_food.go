package biz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/food/model"
)

type ListFoodStorage interface {
	ListFood(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Food, error)
}

type listFoodBiz struct {
	store ListFoodStorage
}

func NewListFoodBiz(store ListFoodStorage) *listFoodBiz {
	return &listFoodBiz{store: store}
}

func (biz *listFoodBiz) ListFood(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.Food, error) {
	data, err := biz.store.ListFood(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil

}
