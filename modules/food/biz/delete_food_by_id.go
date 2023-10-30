package biz

import (
	"context"
	"food-delivery/modules/food/model"
)

type DeleteFoodStorage interface {
	GetFood(ctx context.Context, cond map[string]interface{}) (*model.Food, error)
	DeleteFood(ctx context.Context, cond map[string]interface{}) error
}

type deleteFoodBiz struct {
	store DeleteFoodStorage
}

func NewDeleteFoodBiz(store DeleteFoodStorage) *deleteFoodBiz {
	return &deleteFoodBiz{store: store}
}

func (biz *deleteFoodBiz) DeleteFoodById(ctx context.Context, id int) error {
	data, err := biz.store.GetFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.FoodDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.store.DeleteFood(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
