package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

//InitGraphQLSchema func
func InitGraphQLSchema() graphql.Schema {
	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: queryFields}
	rootMutation := graphql.ObjectConfig{Name: "Mutation", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Printf("failed to create new schema, error: %v", err)
	}

	return schema

}

var queryFields = graphql.Fields{
	"master_users": GetAllUser,
	"master_user":  GetOneUser,
}
var mutationFields = graphql.Fields{
	"create_master_user":      CreateOneUser,
	"update_master_user":      UpdateOneUser,
	"update_many_master_user": UpdateManyUser,
	"delete_master_user":      DeleteUser,
	"delete_many_user":        DeleteManyUser,
}
