package biz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/category/model"
)

type ListCategoryStorage interface {
	ListCategory(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStorage
}

func NewListCategoryBiz(store ListCategoryStorage) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.Category, error) {
	data, err := biz.store.ListCategory(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil

}
