package typedefs

import "github.com/graphql-go/graphql"

//UserSchema type
var UserSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "master_user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
			"created_by": &graphql.Field{
				Type: graphql.String,
			},
			"created_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"modified_by": &graphql.Field{
				Type: graphql.String,
			},
			"modified_date": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

//UserInput for parameter
func UserInput(isIDRequired bool) graphql.FieldConfigArgument {
	if isIDRequired {
		return graphql.FieldConfigArgument{
			"where": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Where Input master_user ID",
						Fields: graphql.InputObjectConfigFieldMap{
							"id": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
					},
				),
			},
		}
	}
	return graphql.FieldConfigArgument{
		"where": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(
				graphql.InputObjectConfig{
					Name: "Where Input master_user Argument",
					Fields: graphql.InputObjectConfigFieldMap{
						//id input
						"id": &graphql.InputObjectFieldConfig{
							Type: graphql.Int,
						},
						"id_lt": &graphql.InputObjectFieldConfig{
							Type: graphql.Int,
						},
						"id_lte": &graphql.InputObjectFieldConfig{
							Type: graphql.Int,
						},
						"id_gt": &graphql.InputObjectFieldConfig{
							Type: graphql.Int,
						},
						"id_gte": &graphql.InputObjectFieldConfig{
							Type: graphql.Int,
						},
						"id_not": &graphql.InputObjectFieldConfig{
							Type: graphql.Int,
						},
						//name input
						"name": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"name_not_contains": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"name_contains": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"name_start_with": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"name_end_with": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						//is_active
						"is_active": &graphql.InputObjectFieldConfig{
							Type: graphql.Boolean,
						},
						//created_by
						"created_by": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"created_by_not_contains": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"created_by_contains": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"created_by_start_with": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"created_by_end_with": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						//created_date
						"created_date": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"created_date_not": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"created_date_lt": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"created_date_lte": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"created_date_gt": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"created_date_gte": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						//modified_by
						"modified_by": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"modified_by_contains": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"modified_by_not_contains": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"modified_by_start_with": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"modified_by_end_with": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						//modified_date
						"modified_date": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"modified_date_not": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"modified_date_lt": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"modified_date_lte": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"modified_date_gt": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
						"modified_date_gte": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime,
						},
					},
				},
			),
		},
	}

}
