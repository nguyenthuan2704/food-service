package ginfood

import (
	"food-delivery/common"
	"food-delivery/modules/food/biz"
	"food-delivery/modules/food/model"
	"food-delivery/modules/food/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateFood(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.FoodCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewCreateFoodBiz(store)
		if err := business.CreateNewFood(c.Request.Context(), &data); err != nil {
			/*c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})*/
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
