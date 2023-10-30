package ginfood

import (
	"food-delivery/common"
	"food-delivery/modules/food/biz"
	"food-delivery/modules/food/model"
	"food-delivery/modules/food/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateFood(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.FoodUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateFoodBiz(store)
		if err := business.UpdateFoodById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
