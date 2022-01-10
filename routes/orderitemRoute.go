package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func OrderItemRoutes (incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItems_id", controllers.GetOrderItem())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateorderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())

}