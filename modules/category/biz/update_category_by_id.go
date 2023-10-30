package biz

import (
	"context"
	"food-delivery/modules/category/model"
)

type UpdateCategoryStorage interface {
	GetFood(ctx context.Context, cond map[string]interface{}) (*model.Category, error)
	UpdateCategory(ctx context.Context, cond map[string]interface{}, dataUpdate *model.CategoryUpdate) error
}

type updateCategoryBiz struct {
	store UpdateCategoryStorage
}

func NewUpdateCategoryBiz(store UpdateCategoryStorage) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategoryById(ctx context.Context, id int, dataUpdate *model.CategoryUpdate) error {
	data, err := biz.store.GetFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		/*return model.ErrItemDeleted*/
		return biz.store.UpdateCategory(ctx, map[string]interface{}{"id": id}, dataUpdate)
	}

	if err := biz.store.UpdateCategory(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
