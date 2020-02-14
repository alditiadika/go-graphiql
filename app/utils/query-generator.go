package utils

import (
	"fmt"
)

//NumericGenerator function
func NumericGenerator(field string, value interface{}, operator string) string {

	switch operator {
	case "_lt":
		return fmt.Sprintf(`and "%s" < %v `, field, value)
	case "_lte":
		return fmt.Sprintf(`and "%s" <= %v `, field, value)
	case "_gt":
		return fmt.Sprintf(`and "%s" > %v `, field, value)
	case "_gte":
		return fmt.Sprintf(`and "%s" >= %v `, field, value)
	case "_not":
		return fmt.Sprintf(`and "%s" <> %v `, field, value)
	default:
		return fmt.Sprintf(`and "%s" = %v `, field, value)
	}
}

//StringGenerator function
func StringGenerator(field string, value interface{}, operator string) string {
	switch operator {
	case "_contains":
		return fmt.Sprintf(`and "%s" like '%s' `, field, stringContatter("both", value))
	case "_not_contains":
		return fmt.Sprintf(`and "%s" not like '%s' `, field, stringContatter("both", value))
	case "_end_with":
		return fmt.Sprintf(`and "%s" like '%s' `, field, stringContatter("before", value))
	case "_start_with":
		return fmt.Sprintf(`and "%s" like '%s' `, field, stringContatter("after", value))
	default:
		return fmt.Sprintf(`and "%s" = '%v' `, field, value)
	}
}

//DateGenerator function
// func DateGenerator(field string, value interface{}, operator string) string {
// 	a, b = time.Parse(databasefield.TimestampzFormat, value.(string))
// 	return a
// }
func stringContatter(condition string, value interface{}) string {
	switch condition {
	case "before":
		return "%" + fmt.Sprintf("%v", value)
	case "after":
		return fmt.Sprintf("%v", value) + "%"
	case "both":
		return "%" + fmt.Sprintf("%v", value) + "%"
	default:
		return ""
	}
}
