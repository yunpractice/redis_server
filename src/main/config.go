package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port  string `json:"Port"`  //监听端口，默认9999
	DbNum int    `json:"DbNum"` //数据库个数，默认16个，序号0-15
}

var config *Config = &Config{
	Port:  "6379",
	DbNum: 16,
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
