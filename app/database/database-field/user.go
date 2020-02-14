package databasefield

//UserFieldType func
func UserFieldType(field string) string {
	switch field {
	case "id":
		return "numeric"
	case "name":
		return "string"
	case "is_active":
		return "boolean"
	case "created_date":
		return "date"
	case "created_by":
		return "string"
	case "modified_date":
		return "date"
	case "modified_by":
		return "string"
	default:
		return ""
	}
}

//UserFieldList function
func UserFieldList() map[int]string {
	a := map[int]string{
		1: "id",
		2: "name",
		3: "is_active",
		4: "created_date",
		5: "created_by",
		6: "modified_date",
		7: "modified_by",
	}
	return a
}
