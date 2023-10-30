package biz

import (
	"context"
	"food-delivery/modules/category/model"
)

type DeleteCategoryStorage interface {
	GetFood(ctx context.Context, cond map[string]interface{}) (*model.Category, error)
	DeleteCategory(ctx context.Context, cond map[string]interface{}) error
}

type deleteCategoryBiz struct {
	store DeleteCategoryStorage
}

func NewDeleteCategoryBiz(store DeleteCategoryStorage) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategoryById(ctx context.Context, id int) error {
	data, err := biz.store.GetFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.store.DeleteCategory(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
