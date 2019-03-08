package utils

import (

)

func ArgsTag(args[]string)string{
	var final string
	for index, tag := range args{
		if index == len(args) -1 {
			final += tag
		} else {
			final += tag+"_"
		}
	}

	return final 
}