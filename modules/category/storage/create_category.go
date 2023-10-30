package storage

import (
	"context"
	"food-delivery/modules/category/model"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *model.CategoryCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
