package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/modules/restaurant/biz"
	"food-delivery/modules/restaurant/model"
	"food-delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurant(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.RestaurantCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewCreateRestaurantBiz(store)
		if err := business.CreateNewRestaurant(c.Request.Context(), &data); err != nil {
			/*c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})*/
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
