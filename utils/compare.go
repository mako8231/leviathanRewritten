package utils 

import (
	"strings"
)

func Compare(str, substr string) bool{
	str, substr = strings.ToLower(str), strings.ToLower(substr)
	return str == substr
}