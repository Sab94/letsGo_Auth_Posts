package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/posts-api/database"
	"github.com/posts-api/types"
	"go.mongodb.org/mongo-driver/bson"
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
