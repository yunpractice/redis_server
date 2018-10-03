package main

import (
	"strings"
)

type CMD func(db *DB, args []string) string

var cmds map[string]CMD = map[string]CMD{}

func register_cmds() {
	register_string_cmds()
	register_list_cmds()
	register_set_cmds()
}

func dispatch_cmd(db *DB, cmd string, args []string) string {
	cmd = strings.ToLower(cmd)
	f, ok := cmds[cmd]
	if !ok {
		return get_error_string("ERR unknown command: " + cmd)
	}
	return f(db, args)
}
