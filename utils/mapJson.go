package utils

import (
	"encoding/json"
)

func MapJSON(body []byte) (map[string]json.RawMessage, error){
	var mapJSON map[string]json.RawMessage
	//Simply unmarshal the body into a map
	err := json.Unmarshal(body,&mapJSON)
	if err != nil{
		return nil, err 
	}
	//And returns this same map
	return mapJSON, nil 	
}