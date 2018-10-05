package main

import (
	"strings"
)

// 命令调用函数，返回结果和、否保存aof日志记录
type CMD func(db *DB, args []string) (string, bool)

var cmds map[string]CMD = map[string]CMD{}

func register_cmds() {
	register_string_cmds()
	register_list_cmds()
	register_set_cmds()
}

func dispatch_cmd(db *DB, cmd string, args []string) (string, bool) {
	cmd = strings.ToLower(cmd)
	f, ok := cmds[cmd]
	if !ok {
		return get_error_string("ERR unknown command: " + cmd), false
	}
	return f(db, args)
}
