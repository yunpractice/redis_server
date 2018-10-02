package main

import (
	"fmt"
)

type CMD func(db *DB, args []string) string

var cmds map[string]CMD = map[string]CMD{}

func register_cmds() {
	register_string_cmds()
}

func dispatch_cmd(db *DB, cmd string, args []string) string {
	f, ok := cmds[cmd]
	if !ok {
		fmt.Println("cmd: ", cmd, " not exist")
		return "-command not exist!\r\n"
	}
	return f(db, args)
}
