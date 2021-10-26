package routes

import (
	"shop/services/v1/items"
	"shop/services/v1/taxes"

	"github.com/gin-gonic/gin"
)

func V1Application(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		r_items := v1.Group("items")
		{
			r_items.GET("/", items.ItemsHandler)
			r_items.POST("/store", items.StoreHandler)
			r_items.PUT("/:id/update", items.UpdateHandler)
			r_items.DELETE("/:id/delete", items.DeleteHandler)
		}
		r_taxes := v1.Group("taxes")
		{
			r_taxes.POST("/store", taxes.StoreHandler)
			r_taxes.PUT("/:id/update", taxes.UpdateHandler)
			r_taxes.DELETE("/:id/delete", taxes.DeleteHandler)
		}
	}
}
