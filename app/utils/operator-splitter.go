package utils

import "strings"

//OperatorSplitter function
func OperatorSplitter(fieldList map[int]string, key string) (string, string) {
	operator := ""
	for _, field := range fieldList {
		splitter := strings.Split(key, field)
		if len(splitter) > 1 {
			return field, splitter[1]
		}
	}
	return key, operator
}
