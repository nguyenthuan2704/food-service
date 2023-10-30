package biz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/food/model"
	"strings"
)

type CreateFoodStorage interface {
	CreateFood(ctx context.Context, data *model.FoodCreation) error
}

type createFoodBiz struct {
	store CreateFoodStorage
}

func NewCreateFoodBiz(store CreateFoodStorage) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateNewFood(ctx context.Context, data *model.FoodCreation) error {
	if strings.TrimSpace(data.Name) == "" {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.CreateFood(ctx, data); err != nil {
		/*return err*/
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
