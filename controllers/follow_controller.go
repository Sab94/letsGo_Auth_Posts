package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/posts-api/database"
	"github.com/posts-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/oauth2.v3/models"
	"log"
)

func Follow (c *gin.Context) {
	followUserId , _ := primitive.ObjectIDFromHex(c.Param("id"))
	collection := database.DB.Collection("users")
	ti, _ := c.Get(ginserver.DefaultConfig.TokenKey)
	token := ti.(*models.Token)

	followUser := types.User{}
	err := collection.FindOne(context.Background(), bson.M{"_id": followUserId}).Decode(&followUser)

	_, err = collection.UpdateOne(context.Background(), bson.M{"email": token.ClientID},
		bson.M{ "$push": bson.M{"following": followUser} })


	if err != nil {
		log.Println(err)
		c.Abort()
	}

	c.String(200, "Followed")
	c.Done()
}

func Unfollow (c *gin.Context) {
	followUserId , _ := primitive.ObjectIDFromHex(c.Param("id"))
	collection := database.DB.Collection("users")
	ti, _ := c.Get(ginserver.DefaultConfig.TokenKey)
	token := ti.(*models.Token)

	followUser := types.User{}
	err := collection.FindOne(context.Background(), bson.M{"_id": followUserId}).Decode(&followUser)

	a, err := collection.UpdateOne(context.Background(), bson.M{"email": token.ClientID},
		bson.M{ "$pull": bson.M{"following": followUser} })


	if err != nil {
		log.Println(err)
		c.Abort()
	}

	fmt.Printf("%+v",a)

	c.String(200, "Unfollowed")
	c.Done()
}