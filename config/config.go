package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Token       string `json:"token"`
	Prefix      string `json:"prefix"`
	Owner       string `json:"owner"`
	DefaultPort string `json:"default_port"`
	SafeMode    bool   `json:"safe_mode"`
	Server      string `json:"server"`
	PoolChan    string `json:"pool_chan"`
}

var (
	Data       *Config
	BotVersion string
)

func LoadConfig() {
	//Load the .json file
	b, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	//Parsing the .json file
	err = json.Unmarshal(b, &Data)
	if err != nil {
		log.Fatal(err.Error())
	}

	port := os.Getenv("PORT")
	if port != "" {
		Data.DefaultPort = port
	}

	BotVersion = "1.2"
}
