package middlewares

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

type AuthObj struct {
	ClientID   string      `json:"ClientID"`
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
		token := ti.(*models.Token)
		if exists {
			a := getUserId(token.ClientID)
			fmt.Printf("%+v",a)
			token.UserID = a.Hex()
			c.Set(ginserver.DefaultConfig.TokenKey, token)
			return
		}

		c.Next()
	}
}


func getUserId (email string) primitive.ObjectID {
	collection := database.DB.Collection("users")

	user := types.User{}
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		log.Println(err)
	}

	return user.Id

}
