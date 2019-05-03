package fields

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/posts-api/database"
	"github.com/posts-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Users = &graphql.Field{
	Type: graphql.NewList(types.UserGQLObj),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userCollection := database.DB.Collection("users")
		_options := options.FindOptions{}
		ctx := context.Background()
		// Sort by `_id` field descending
		_options.Sort = bson.D{{"_id", -1}}
		users, err := userCollection.Find(ctx, bson.D{}, &_options)
		if err != nil {
			log.Fatal(err)
		}
		usersGql:= []types.User{}
		for users.Next(context.Background()) {
			result := types.User{}
			err := users.Decode(&result)
			if err != nil { panic(err) }

			var inInterface map[string]interface{}
			inrec, _ := json.Marshal(result)
			json.Unmarshal(inrec, &inInterface)


			// convert BSON to struct
			user := types.User{}
			for key, value := range inInterface {

				switch (key) {
				case "name":
					user.Name = fmt.Sprintf("%v",value)
				case "email":
					user.Email= fmt.Sprintf("%v",value)
				case "id":
					user.Id, _= primitive.ObjectIDFromHex(fmt.Sprintf("%v",value))
				default:
				}
			}
			usersGql = append(usersGql, user)
		}

		return usersGql, nil
	},
}
