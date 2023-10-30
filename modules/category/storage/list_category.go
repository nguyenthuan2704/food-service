package storage

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/category/model"
)

func (s *sqlStore) ListCategory(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Category, error) {
	var result []model.Category
	db := s.db.Where("status <> ?", "Deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.Category{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
