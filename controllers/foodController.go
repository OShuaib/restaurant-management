package controllers

import (
	"context"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client,"food")
var validate = validator.New()

func GetFood() gin.HandlerFunc{
	return func (c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id":foodId}).Decode(&food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food"})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func GetFoods() gin.HandlerFunc{
	return func (c *gin.Context){

	}
}

func CreateFood() gin.HandlerFunc{
	return func (c *gin.Context){
		var ctx, cancel= context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err := menuCollection.FindOne(ctx,bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu was not found"})
			return
		}

		food.Created_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Update_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(food.Price,2)
		food.Price = num


		result, insertErr := foodCollection.InsertOne(ctx,food)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Food item was not created"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}


func round (num float64) int {

}

func toFixed(num float64, precision int) float64{

}

func UpdateFood() gin.HandlerFunc{
	return func (c *gin.Context){
		
	}
}



