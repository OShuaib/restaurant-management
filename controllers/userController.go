package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")

func GetUsers() gin.HandlerFunc{
	return func (c *gin.Context)  {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPage< 1 {
			recordPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1{
			page = 1
		}

		startIndex := (page - 1) * recordPage
		startIndex ,err = strconv.Atoi(c.Query("startIndex"))


		matchStage := bson.D{{"$match", bson.D{{}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_item", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPage}}}},
			}}}
			result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
				matchStage, projectStage,
			})
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
				return
			}

			var allUsers []bson.M
			if err = result.All(ctx,&allUsers); err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, allUsers[0])
	}
}

func GetUser() gin.HandlerFunc {
	return func (c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		userId := c.Param("user_id")

		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error", "error occured while listing user items"})
		}

		c.JSON(http.StatusOK, user)
	}
}

func SignUp() gin.HandlerFunc {
	return func (c *gin.Context){
		//convert the JSON data coming from postman to something that golang understands

		//validate the data based on user struct

		//check if the email has already been used by another user

		//hashpassword

		//check if the phone no. has already been used by another user

		//create some extra details for the user object - created_at, updated_at, ID

		//generate token and refresh token (generate all token from helper)

		//if all check passed, insert the new user into the user collection

		//return status OK and send result back

	}
}

func Login() gin.HandlerFunc{
	return func (c *gin.Context){
		//convert the login JSON data from postman to what golang understand

		//find a user with that email and see if that user even exists

		//verify password

		//if all goes well, then we generate tokens

		//update tokens - token and refresh token

		//return statusOK
	}
}


func HashPassword(password string) string {

	return ""
}

func VerifyPassword (userPassword string, providePassword string)(bool, string){
	
	return true,""
}
