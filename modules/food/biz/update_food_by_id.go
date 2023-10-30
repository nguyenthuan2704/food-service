package biz

import (
	"context"
	"food-delivery/modules/food/model"
)

type UpdateFoodStorage interface {
	GetFood(ctx context.Context, cond map[string]interface{}) (*model.Food, error)
	UpdateFood(ctx context.Context, cond map[string]interface{}, dataUpdate *model.FoodUpdate) error
}

type updateFoodBiz struct {
	store UpdateFoodStorage
}

func NewUpdateFoodBiz(store UpdateFoodStorage) *updateFoodBiz {
	return &updateFoodBiz{store: store}
}

func (biz *updateFoodBiz) UpdateFoodById(ctx context.Context, id int, dataUpdate *model.FoodUpdate) error {
	data, err := biz.store.GetFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.FoodDeleted {
		/*return model.ErrItemDeleted*/
		return biz.store.UpdateFood(ctx, map[string]interface{}{"id": id}, dataUpdate)
	}

	if err := biz.store.UpdateFood(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
