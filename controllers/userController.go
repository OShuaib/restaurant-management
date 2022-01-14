package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurant-management/database"
	helper "restaurant-management/helpers"
	"restaurant-management/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		//convert the JSON data coming from postman to something that golang understands
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//validate the data based on user struct
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		//check if the email has already been used by another user
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err !=nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
			return
		}

		//hashpassword
		password := HashPassword(*&user.Password)
		user.Password = password

		//check if the phone no. has already been used by another user
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err !=nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the phone"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number already exists"})
			return 
		}

		//create some extra details for the user object - created_at, updated_at, ID
		user.Created_at,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		user.Updated_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		//generate token and refresh token (generate all token from helper)
		token , refreshToken,_ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		user.Token = &token
		user.Refresh_Token = &refreshToken 

		//if all check passed, insert the new user into the user collection
		resultInsertion, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"USER ITEM WAS NOT CREATED"})
			return 
		}
		defer cancel()

		//return status OK and send result back
		c.JSON(http.StatusOK, resultInsertion)

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
