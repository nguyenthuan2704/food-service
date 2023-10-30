package main

import (
	"fmt"
	"food-delivery/common"
	"food-delivery/middleware"
	gincategory "food-delivery/modules/category/transport/gin"
	ginfood "food-delivery/modules/food/transport/gin"
	ginrestaurant "food-delivery/modules/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	v1 := r.Group("/v1", middleware.Recovery())
	{
		categories := v1.Group("/categories")
		{
			categories.GET("", gincategory.ListCategory(db))
			categories.GET("/:id", gincategory.GetCategory(db))
			categories.POST("", gincategory.CreateCategory(db))
			categories.PATCH("/:id", gincategory.UpdateItem(db))
			categories.DELETE("/:id", gincategory.DeleteCategory(db))
		}
		foods := v1.Group("/foods")
		{
			foods.GET("", ginfood.ListFood(db))
			foods.GET("/:id", ginfood.GetFood(db))
			foods.POST("", ginfood.CreateFood(db))
			foods.PATCH("/:id", ginfood.UpdateFood(db))
			foods.DELETE("/:id", ginfood.DeleteFood(db))
		}
		restaurants := v1.Group("/restaurants")
		{
			restaurants.GET("", ginrestaurant.ListRestaurant(db))
			restaurants.GET("/:id", ginrestaurant.GetRestaurant(db))
			restaurants.POST("", ginrestaurant.CreateRestaurant(db))
			restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(db))
			restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3009")
}
