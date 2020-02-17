package utils

import (
	"fmt"
	"strings"
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
	vstr := value.(string)
	switch operator {
	case "_contains":
		return fmt.Sprintf(`and "%s" like '%s' `, field, stringContatter("both", strings.ToUpper(vstr)))
	case "_not_contains":
		return fmt.Sprintf(`and "%s" not like '%s' `, field, stringContatter("both", strings.ToUpper(vstr)))
	case "_end_with":
		return fmt.Sprintf(`and "%s" like '%s' `, field, stringContatter("before", strings.ToUpper(vstr)))
	case "_start_with":
		return fmt.Sprintf(`and "%s" like '%s' `, field, stringContatter("after", strings.ToUpper(vstr)))
	default:
		return fmt.Sprintf(`and "%s" = '%v' `, field, strings.ToUpper(vstr))
	}
}

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
