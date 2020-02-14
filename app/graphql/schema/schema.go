package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

//InitGraphQLSchema func
func InitGraphQLSchema() graphql.Schema {
	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: queryFields}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
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
var mutationFields = graphql.Fields{}
