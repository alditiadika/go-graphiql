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
	Args:        typedefs.GetUserInput(false, "Where Input get many master_user ID"),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.GetUser(where), nil
	},
}

//GetOneUser field
var GetOneUser = &graphql.Field{
	Type:        typedefs.UserSchema,
	Description: "Get One User",
	Args:        typedefs.GetUserInput(true, "Where Input get one master_user ID"),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.GetOneUser(where), nil
	},
}

//CreateOneUser field
var CreateOneUser = &graphql.Field{
	Type:        typedefs.UserSchema,
	Description: "Create One User",
	Args:        typedefs.CreateUserInput,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		data, _ := params.Args["data"].(map[string]interface{})
		return resolver.CreateUser(data), nil
	},
}

//UpdateOneUser field
var UpdateOneUser = &graphql.Field{
	Type:        typedefs.UserSchema,
	Description: "Update One User",
	Args:        typedefs.UpdateOneUserInput,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		data, _ := params.Args["data"].(map[string]interface{})
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.UpdateOneUser(data, where), nil
	},
}

//UpdateManyUser field
var UpdateManyUser = &graphql.Field{
	Type:        typedefs.ManyDataAffected,
	Description: "Update Many User",
	Args:        typedefs.UpdateManyUserInput,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		data, _ := params.Args["data"].(map[string]interface{})
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.UpdateManyUser(data, where), nil
	},
}

//DeleteUser field
var DeleteUser = &graphql.Field{
	Type:        typedefs.UserSchema,
	Description: "Delete One User",
	Args:        typedefs.GetUserInput(true, "Where Input delete one master_user ID"),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.DeleteUser(where), nil
	},
}

//DeleteManyUser field
var DeleteManyUser = &graphql.Field{
	Type:        typedefs.ManyDataAffected,
	Description: "Delete Many User",
	Args:        typedefs.GetUserInput(false, "Where Input delete many master_user ID"),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		where, _ := params.Args["where"].(map[string]interface{})
		return resolver.DeleteManyUser(where), nil
	},
}
