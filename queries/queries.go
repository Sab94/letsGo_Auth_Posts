package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/posts-api/queries/fields"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"hello": fields.Hello,
		"world": fields.World,
		"users": fields.Users,
	},
})