package storage

import (
	"context"
	"food-delivery/modules/restaurant/model"
)

func (s *sqlStore) UpdateRestaurant(ctx context.Context, cond map[string]interface{}, dataUpdate *model.RestaurantUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
