package utils

import "net/url"

func Encode(str string) string {
	u, err := url.Parse(str)
	if err != nil{
		return ""
	}

	return u.String()
}