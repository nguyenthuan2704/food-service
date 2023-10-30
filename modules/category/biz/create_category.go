package biz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/category/model"
	"strings"
)

type CreateCategoryStorage interface {
	CreateCategory(ctx context.Context, data *model.CategoryCreation) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateNewCategory(ctx context.Context, data *model.CategoryCreation) error {
	if strings.TrimSpace(data.Name) == "" {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.CreateCategory(ctx, data); err != nil {
		/*return err*/
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
