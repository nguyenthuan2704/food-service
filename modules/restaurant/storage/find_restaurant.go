package storage

import (
	"context"
	"food-delivery/modules/restaurant/model"
)

func (s *sqlStore) GetRestaurant(ctx context.Context, cond map[string]interface{}) (*model.Restaurant, error) {
	var data model.Restaurant

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
