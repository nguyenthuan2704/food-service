package biz

import (
	"context"
	"food-delivery/modules/food/model"
)

type GetFoodStorage interface {
	GetFood(ctx context.Context, cond map[string]interface{}) (*model.Food, error)
}

type getFoodBiz struct {
	store GetFoodStorage
}

func NewGetFoodBiz(store GetFoodStorage) *getFoodBiz {
	return &getFoodBiz{store: store}
}

func (biz *getFoodBiz) GetFoodById(ctx context.Context, id int) (*model.Food, error) {
	data, err := biz.store.GetFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
