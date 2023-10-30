package storage

import (
	"context"
	"food-delivery/modules/food/model"
)

func (s *sqlStore) UpdateFood(ctx context.Context, cond map[string]interface{}, dataUpdate *model.FoodUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
