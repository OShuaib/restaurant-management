package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc{
	return func (c *gin.Context){
		
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc{
	return func (c *gin.Context){
		
	}
}

func ItemsByOder(id string)(OrderItems []primitive.M, err error) {
	
		
}

func GetOrderItem() gin.HandlerFunc{
	return func (c *gin.Context){
		
	}
}

func UpdateOrderItem() gin.HandlerFunc{
	return func (c *gin.Context){
		
	}
}

func CreateOrderItem() gin.HandlerFunc{
	return func (c *gin.Context){
		
	}
}