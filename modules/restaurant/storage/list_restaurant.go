package storage

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/model"
)

func (s *sqlStore) ListRestaurant(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	var result []model.Restaurant
	db := s.db.Where("status <> ?", "Deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Order("id asc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
