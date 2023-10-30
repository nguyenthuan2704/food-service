package storage

import (
	"context"
	"food-delivery/modules/food/model"
)

func (s *sqlStore) DeleteFood(ctx context.Context, cond map[string]interface{}) error {
	/*Hard Delete - Xóa thẳng record */
	/*	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}*/
	/*Soft Delete - Update lại record của column status thành Deleted*/
	deletedStatus := model.FoodDeleted
	if err := s.db.Table(model.Food{}.TableName()).Where(cond).Updates(map[string]interface{}{
		//"status: "Deleted",
		"status": deletedStatus.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}
