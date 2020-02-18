package typedefs

import (
	"time"

	databasefield "github.com/alditiadika/go-graphiql/app/database/database-field"
	"github.com/graphql-go/graphql"
)

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

//GetUserInput for parameter
func GetUserInput(isIDRequired bool, desc string) graphql.FieldConfigArgument {
	if isIDRequired {
		return graphql.FieldConfigArgument{
			"where": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: desc,
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
					Name: desc,
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

//CreateUserInput function
var CreateUserInput graphql.FieldConfigArgument = graphql.FieldConfigArgument{
	"data": &graphql.ArgumentConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "Data Input create master_user Argument",
				Fields: graphql.InputObjectConfigFieldMap{
					"name": &graphql.InputObjectFieldConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"is_active": &graphql.InputObjectFieldConfig{
						Type:         graphql.Boolean,
						DefaultValue: false,
					},
					"created_by": &graphql.InputObjectFieldConfig{
						Type:         graphql.String,
						DefaultValue: "ADMIN",
					},
					"created_date": &graphql.InputObjectFieldConfig{
						Type:         graphql.DateTime,
						DefaultValue: time.Now().Format(databasefield.TimestampzFormat),
					},
					"modified_by": &graphql.InputObjectFieldConfig{
						Type:         graphql.String,
						DefaultValue: "ADMIN",
					},
					"modified_date": &graphql.InputObjectFieldConfig{
						Type:         graphql.DateTime,
						DefaultValue: time.Now().Format(databasefield.TimestampzFormat),
					},
				},
			},
		),
	},
}

//UpdateOneUserInput function
var UpdateOneUserInput graphql.FieldConfigArgument = graphql.FieldConfigArgument{
	"where": &graphql.ArgumentConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "Where Input update master_user ID",
				Fields: graphql.InputObjectConfigFieldMap{
					"id": &graphql.InputObjectFieldConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
			},
		),
	},
	"data": &graphql.ArgumentConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "Data Input update master_user Argument",
				Fields: graphql.InputObjectConfigFieldMap{
					"name": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"is_active": &graphql.InputObjectFieldConfig{
						Type: graphql.Boolean,
					},
					"created_by": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"created_date": &graphql.InputObjectFieldConfig{
						Type: graphql.DateTime,
					},
					"modified_by": &graphql.InputObjectFieldConfig{
						Type:         graphql.String,
						DefaultValue: "ADMIN",
					},
					"modified_date": &graphql.InputObjectFieldConfig{
						Type:         graphql.DateTime,
						DefaultValue: time.Now().Format(databasefield.TimestampzFormat),
					},
				},
			},
		),
	},
}

//UpdateManyUserInput var
var UpdateManyUserInput graphql.FieldConfigArgument = graphql.FieldConfigArgument{
	"where": &graphql.ArgumentConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "Where Input update many master_user Argument",
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
	"data": &graphql.ArgumentConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "Data Input update many master_user Argument",
				Fields: graphql.InputObjectConfigFieldMap{
					"name": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"is_active": &graphql.InputObjectFieldConfig{
						Type: graphql.Boolean,
					},
					"created_by": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"created_date": &graphql.InputObjectFieldConfig{
						Type: graphql.DateTime,
					},
					"modified_by": &graphql.InputObjectFieldConfig{
						Type:         graphql.String,
						DefaultValue: "ADMIN",
					},
					"modified_date": &graphql.InputObjectFieldConfig{
						Type:         graphql.DateTime,
						DefaultValue: time.Now().Format(databasefield.TimestampzFormat),
					},
				},
			},
		),
	},
}
