package gincategory

import (
	"food-delivery/common"
	"food-delivery/modules/category/biz"
	"food-delivery/modules/category/model"
	"food-delivery/modules/category/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CategoryCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewCreateCategoryBiz(store)
		if err := business.CreateNewCategory(c.Request.Context(), &data); err != nil {
			/*c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})*/
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
