package fields

import "github.com/graphql-go/graphql"

var World = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "hello", nil
	},
}
