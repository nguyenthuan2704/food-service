package storage

import (
	"context"
	"food-delivery/modules/food/model"
)

func (s *sqlStore) GetFood(ctx context.Context, cond map[string]interface{}) (*model.Food, error) {
	var data model.Food

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
