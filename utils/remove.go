package utils

import (
	"strings"
)

func Remove(str string, s string) string {
	str = strings.Replace(str, s, "", -1)
	return str 
}