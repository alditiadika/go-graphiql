package schema

import (
	"github.com/alditiadika/go-graphiql/app/graphql/resolver"
	typedefs "github.com/alditiadika/go-graphiql/app/graphql/type-defs"
	"github.com/graphql-go/graphql"
)

//GetAllUser field
var GetAllUser = &graphql.Field{
	Type:        graphql.NewList(typedefs.UserSchema),
	Description: "Get All Users",
	Args:        typedefs.UserInput(false),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.GetUser(where), nil
	},
}

//GetOneUser field
var GetOneUser = &graphql.Field{
	Type:        graphql.NewList(typedefs.UserSchema),
	Description: "Get One User",
	Args:        typedefs.UserInput(true),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.GetUser(where), nil
	},
}
