package storage

import (
	"context"
	"food-delivery/modules/category/model"
)

func (s *sqlStore) UpdateCategory(ctx context.Context, cond map[string]interface{}, dataUpdate *model.CategoryUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
