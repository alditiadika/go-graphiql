package generator

import (
	"fmt"

	dbField "github.com/alditiadika/go-graphiql/app/database/database-field"
	"github.com/alditiadika/go-graphiql/app/utils"
)

//GetUserByParam func
func GetUserByParam(args map[string]interface{}) string {
	fieldList := dbField.UserFieldList()
	qs := "select * from master_user mu where 1=1 "

	for key, value := range args {
		field, operator := utils.OperatorSplitter(fieldList, key)
		fieldType := dbField.UserFieldType(field)

		switch fieldType {
		case "numeric":
			qs += utils.NumericGenerator(field, value, operator)
			break
		case "string":
			qs += utils.StringGenerator(field, value, operator)
			break
		case "boolean":
			qs += fmt.Sprintf(`and "%s" = %v `, field, value)
			break
		default:
			qs += ""
			break
		}
	}
	return qs
}
