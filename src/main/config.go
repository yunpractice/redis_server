package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port      string `json:"Port"`      //监听端口，默认9999
	DbNum     int    `json:"DbNum"`     //数据库个数，默认16个，序号0-15
	AsyncMode string `json:"ASYNCMODE"` //aof保存模式 AOF_FSYNC_NO ：不保存 AOF_FSYNC_EVERYSEC ：每一秒钟保存一次 AOF_FSYNC_ALWAYS ：每执行一个命令保存一次
	AsyncFile string `json:"ASYNCFILE"` //aof文件
}

var config *Config = &Config{
	Port:      "6379",
	DbNum:     16,
	AsyncMode: "AOF_FSYNC_NO",
	AsyncFile: "./aof.log",
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
