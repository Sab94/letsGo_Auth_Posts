package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/posts-api/database"
	"github.com/posts-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/oauth2.v3/models"
	"log"
)

func WhoAmI(c *gin.Context) {
	ti, _ := c.Get(ginserver.DefaultConfig.TokenKey)
	token := ti.(*models.Token)

	userCollection := database.DB.Collection("users")
	user := types.User{}
	err := userCollection.FindOne(context.Background(), bson.M{"email": token.ClientID}).Decode(&user)

	if err != nil {
		log.Println(err)
		c.Abort()
	}

	c.JSON(200, user)
	c.Done()
}


func GetAllUsers(c *gin.Context) {
	userCollection := database.DB.Collection("users")
	_options := options.FindOptions{}
	ctx := context.Background()
	// Sort by `_id` field descending
	_options.Sort = bson.D{{"_id", -1}}

	users:= []types.User{}
	cur, err := userCollection.Find(ctx, bson.D{}, &_options)

	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		result := types.User{}
		err := cur.Decode(&result)
		users = append(users, result)
		if err != nil { log.Fatal(err) }
	}

	if err != nil {
		log.Println(err)
		c.Abort()
	}

	c.JSON(200, users)
	c.Done()
}