package utils

import (
	"strings"
)

func GetArgs(str, rep string) []string {
	//Remove the command name
	str = strings.Replace(str, rep, "", 1)

	return strings.Fields(str)
}
