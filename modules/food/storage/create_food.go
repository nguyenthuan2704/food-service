package storage

import (
	"context"
	"food-delivery/modules/food/model"
)

func (s *sqlStore) CreateFood(ctx context.Context, data *model.FoodCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
