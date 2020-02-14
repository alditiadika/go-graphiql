package resolver

import (
	"fmt"

	"github.com/alditiadika/go-graphiql/app/database"
	"github.com/alditiadika/go-graphiql/app/database/generator"
	"github.com/alditiadika/go-graphiql/app/models"
)

//GetUser function
func GetUser(args map[string]interface{}) []models.UserModel {
	q := generator.GetUserByParam(args)

	rows, err := database.RunQuery(q)
	if err != nil {
		fmt.Println("error when run query", err)
		return nil
	}
	var users []models.UserModel
	for rows.Next() {
		item := models.UserModel{}
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.IsActive,
			&item.CreatedBy,
			&item.CreatedDate,
			&item.ModifiedBy,
			&item.ModifiedDate,
		)
		users = append(users, item)
	}
	return users
}
