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
	"time"
)

func Post (c *gin.Context) {
	post := types.Post{}
	ctx := context.Background()
	collection := database.DB.Collection("post")
	err := c.BindJSON(&post)
	ti, _ := c.Get(ginserver.DefaultConfig.TokenKey)
	token := ti.(*models.Token)

	userCollection := database.DB.Collection("users")
	user := types.User{}
	err = userCollection.FindOne(context.Background(), bson.M{"email": token.ClientID}).Decode(&user)
	fmt.Printf("%+v",token)
	fmt.Printf("%+v",user)

	if err != nil {
		log.Println(err)
		c.Abort()
	}

	post.Poster = user
	post.Id = primitive.NewObjectID()
	post.PostedAt = time.Now().String()

	_, err = collection.InsertOne(ctx, post)
	if err != nil {
		log.Println(err)
		c.Abort()
	}
	c.JSON(200, post)
	c.Done()
}

func Posts(c *gin.Context) {
	last100 := []types.Post{}
	ctx := context.Background()
	collection := database.DB.Collection("post")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		result := types.Post{}
		err := cur.Decode(&result)
		last100 = append(last100, result)
		if err != nil { log.Fatal(err) }
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, last100)
	c.Done()
}

