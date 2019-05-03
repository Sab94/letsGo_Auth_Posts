package types

import (
	"github.com/graphql-go/graphql"
	"github.com/posts-api/scalars"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	Email   string `json:"email"`
	Password        string `json:"password"`
	Following []*User `json:"following"`
}

type LoginRequest struct {
	Email   string `json:"email"`
	Password        string `json:"password"`
}

var UserGQLObj = graphql.NewObject(graphql.ObjectConfig {
	Name: "UserGQL",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: scalars.ObjectId,
		},
	},
})