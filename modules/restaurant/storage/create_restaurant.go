package storage

import (
	"context"
	"food-delivery/modules/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *model.RestaurantCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
