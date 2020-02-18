package generator

import (
	"fmt"
	"strings"

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
		case "date":
			qs += utils.NumericGenerator(field, value, operator)
		default:
			qs += ""
			break
		}
	}
	return qs
}

//CreateUserData func
func CreateUserData(args map[string]interface{}) string {
	mutation := "insert into master_user "
	fieldname := "("
	valueField := "("
	index := 0
	for key, value := range args {
		index++
		fieldType := dbField.UserFieldType(key)
		switch fieldType {
		case "string":
			vstr := value.(string)
			if index == len(args) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`'%v')`, strings.ToUpper(vstr))
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`'%v', `, strings.ToUpper(vstr))
			}
			break
		case "numeric":
			if index == len(args) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`%v)`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`%v, `, value)
			}
			break
		case "boolean":
			if index == len(args) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`%v)`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`%v, `, value)
			}
			break
		case "date":
			if index == len(args) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`'%v')`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`'%v', `, value)
			}
			break
		default:
			break
		}
	}
	return fmt.Sprintf(`%s %s values %s`, mutation, fieldname, valueField)
}

//UpdateOneUser function
func UpdateOneUser(data, cond map[string]interface{}) (string, interface{}) {
	mutation := "update master_user set "
	fieldname := "("
	valueField := "("
	whereCond := "where "
	id := 0
	index := 0
	for key, value := range data {
		index++
		fieldType := dbField.UserFieldType(key)
		switch fieldType {
		case "string":
			vstr := value.(string)
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`'%v')`, strings.ToUpper(vstr))
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`'%v', `, strings.ToUpper(vstr))
			}
			break
		case "numeric":
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`%v)`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`%v, `, value)
			}
			break
		case "boolean":
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`%v)`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`%v, `, value)
			}
			break
		case "date":
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`'%v')`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`'%v', `, value)
			}
			break
		default:
			break
		}
	}
	for key, value := range cond {
		whereCond += fmt.Sprintf(`%s = %v`, key, value)
		id = value.(int)
	}
	return fmt.Sprintf(`%s %s = %s %s`, mutation, fieldname, valueField, whereCond), id
}

//UpdateManyUser function
func UpdateManyUser(data, cond map[string]interface{}) string {
	fieldList := dbField.UserFieldList()
	mutation := "update master_user set "
	fieldname := "("
	valueField := "("
	whereCond := "where 1=1 "
	index := 0
	for key, value := range data {
		index++
		fieldType := dbField.UserFieldType(key)
		switch fieldType {
		case "string":
			vstr := value.(string)
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`'%v')`, strings.ToUpper(vstr))
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`'%v', `, strings.ToUpper(vstr))
			}
			break
		case "numeric":
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`%v)`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`%v, `, value)
			}
			break
		case "boolean":
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`%v)`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`%v, `, value)
			}
			break
		case "date":
			if index == len(data) {
				fieldname += fmt.Sprintf(`%s)`, key)
				valueField += fmt.Sprintf(`'%v')`, value)
			} else {
				fieldname += fmt.Sprintf(`%s, `, key)
				valueField += fmt.Sprintf(`'%v', `, value)
			}
			break
		default:
			break
		}
	}
	for key, value := range cond {
		field, operator := utils.OperatorSplitter(fieldList, key)
		fieldType := dbField.UserFieldType(field)

		switch fieldType {
		case "numeric":
			whereCond += utils.NumericGenerator(field, value, operator)
			break
		case "string":
			whereCond += utils.StringGenerator(field, value, operator)
			break
		case "boolean":
			whereCond += fmt.Sprintf(`and "%s" = %v `, field, value)
			break
		case "date":
			whereCond += utils.NumericGenerator(field, value, operator)
		default:
			whereCond += ""
			break
		}
	}
	return fmt.Sprintf(`%s %s = %s %s`, mutation, fieldname, valueField, whereCond)
}

//DeleteUser function
func DeleteUser(args map[string]interface{}) string {
	qs := "delete from master_user where 1=1 "
	fieldList := dbField.UserFieldList()
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
		case "date":
			qs += utils.NumericGenerator(field, value, operator)
		default:
			qs += ""
			break
		}
	}
	return qs
}
