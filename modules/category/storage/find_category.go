package storage

import (
	"context"
	"food-delivery/modules/category/model"
)

func (s *sqlStore) GetFood(ctx context.Context, cond map[string]interface{}) (*model.Category, error) {
	var data model.Category

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
