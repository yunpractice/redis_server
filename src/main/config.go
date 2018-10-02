package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port string `json:"Port"`
}

var config *Config = &Config{
	Port: "9999",
}

func (config *Config) Load(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("wrong config file!")
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
}
