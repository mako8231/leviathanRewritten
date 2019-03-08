package utils

import (

)

func ArgList(args []string)[]string{
	if len(args) == 0{
		args = append(args, "")
		return args
	}

	return args
}