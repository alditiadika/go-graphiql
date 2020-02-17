package typedefs

import "github.com/graphql-go/graphql"

//RootType struct
type RootType struct {
	Query    string `json:"query"`
	Mutation string `json:"mutation"`
}

//ManyDataAffected struct
var ManyDataAffected = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "many_rows_affected",
		Fields: graphql.Fields{
			"rows_affected": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

//DataAffected struct
type DataAffected struct {
	RowsAffected int64 `json:"rows_affected"`
}
