package routes

import (
	"github.com/MarNawar/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incommingRoutes.GET("/orderItems-order:/order_id", controllers.GetOrderItemsByOrder())
	incommingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}