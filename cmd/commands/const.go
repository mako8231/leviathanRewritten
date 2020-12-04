package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type commands struct {
	List []list `json:"commands"`
}

type list struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	CommandList = loadList()
)

func loadList() *commands {
	var com *commands
	b, err := ioutil.ReadFile("cmd/modules/cmd.json")
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	err = json.Unmarshal(b, &com)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return com
}
