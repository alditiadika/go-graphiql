package resolver

import (
	"fmt"

	"github.com/alditiadika/go-graphiql/app/database"
	"github.com/alditiadika/go-graphiql/app/database/generator"
	typedefs "github.com/alditiadika/go-graphiql/app/graphql/type-defs"
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

//GetOneUser function
func GetOneUser(args map[string]interface{}) interface{} {
	q := generator.GetUserByParam(args)

	rows, err := database.RunQuery(q)
	if err != nil {
		fmt.Println("error when run query", err)
		return models.UserModel{}
	}
	item := models.UserModel{}
	for rows.Next() {
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.IsActive,
			&item.CreatedBy,
			&item.CreatedDate,
			&item.ModifiedBy,
			&item.ModifiedDate,
		)
	}
	if item.ID == 0 {
		return nil
	}
	return item
}

//CreateUser function
func CreateUser(args map[string]interface{}) interface{} {
	q := generator.CreateUserData(args)
	_, err := database.RunMutation(q)
	if err != nil {
		fmt.Println("error when create rows", err)
		fmt.Println(q)
		return err
	}
	rows, err := database.RunQuery(`select * from master_user mu order by "id" desc limit 1`)
	if err != nil {
		fmt.Println("error when run query", err)
		return err
	}
	item := models.UserModel{}
	for rows.Next() {
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.IsActive,
			&item.CreatedBy,
			&item.CreatedDate,
			&item.ModifiedBy,
			&item.ModifiedDate,
		)
	}
	return item
}

//UpdateOneUser function
func UpdateOneUser(data, cond map[string]interface{}) interface{} {
	q, id := generator.UpdateOneUser(data, cond)
	_, err := database.RunMutation(q)
	if err != nil {
		fmt.Println("error when create rows", err)
		fmt.Println(q)
		return err
	}
	qGet := fmt.Sprintf(`select * from master_user mu where id = %v`, id)
	rows, err := database.RunQuery(qGet)
	if err != nil {
		fmt.Println("error when run query", err)
		return err
	}
	item := models.UserModel{}
	for rows.Next() {
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.IsActive,
			&item.CreatedBy,
			&item.CreatedDate,
			&item.ModifiedBy,
			&item.ModifiedDate,
		)
	}
	return item
}

//UpdateManyUser function
func UpdateManyUser(data, where map[string]interface{}) interface{} {
	q := generator.UpdateManyUser(data, where)
	rAffected, err := database.RunMutation(q)
	if rAffected == nil {
		rAffected = 0
	}
	rAff := typedefs.DataAffected{RowsAffected: rAffected.(int64)}
	if err != nil {
		fmt.Println("error when create rows", err)
		fmt.Println(q)
		return err
	}
	return rAff
}

//DeleteUser function
func DeleteUser(args map[string]interface{}) interface{} {
	return nil
}

//DeleteManyUser function
func DeleteManyUser(args map[string]interface{}) interface{} {
	return nil
}
