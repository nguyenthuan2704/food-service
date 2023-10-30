package biz

import (
	"context"
	"food-delivery/modules/category/model"
)

type GetCategoryStorage interface {
	GetFood(ctx context.Context, cond map[string]interface{}) (*model.Category, error)
}

type getCategoryBiz struct {
	store GetCategoryStorage
}

func NewGetItemBiz(store GetCategoryStorage) *getCategoryBiz {
	return &getCategoryBiz{store: store}
}

func (biz *getCategoryBiz) GetCategoryById(ctx context.Context, id int) (*model.Category, error) {
	data, err := biz.store.GetFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
