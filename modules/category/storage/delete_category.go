package storage

import (
	"context"
	"food-delivery/modules/category/model"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, cond map[string]interface{}) error {
	/*Hard Delete - Xóa thẳng record */
	/*	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}*/
	/*Soft Delete - Update lại record của column status thành Deleted*/
	deletedStatus := model.ItemStatusDeleted
	if err := s.db.Table(model.Category{}.TableName()).Where(cond).Updates(map[string]interface{}{
		//"status: "Deleted",
		"status": deletedStatus.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}
