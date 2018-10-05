package main

import (
	"container/list"
)

func cmd_lpop(db *DB, args []string) (string, bool) {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'lpop' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_empty_bulk_string(), false
	} else if o.t != 2 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	l := o.p.(*list.List)
	if l.Len() == 0 {
		return get_empty_bulk_string(), false
	}

	node := l.Front()
	l.Remove(node)

	return get_bulk_string(node.Value.(string)), true
}

func cmd_lpush(db *DB, args []string) (string, bool) {
	if len(args) < 2 {
		return get_error_string("ERR wrong number of arguments for 'lpush' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	var l *list.List
	if o == nil {
		l = list.New()
		db.dict.Set(key, &Object{
			t: 2,
			p: l,
		})
	} else if o.t != 2 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	} else {
		l = o.p.(*list.List)
	}

	for i := 1; i < len(args); i++ {
		l.PushFront(args[i])
	}

	return get_number_string(l.Len()), true
}

func cmd_rpop(db *DB, args []string) (string, bool) {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'rpop' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_empty_bulk_string(), false
	} else if o.t != 2 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	l := o.p.(*list.List)
	if l.Len() == 0 {
		return get_empty_bulk_string(), false
	}

	node := l.Back()
	l.Remove(node)

	return get_bulk_string(node.Value.(string)), true
}

func cmd_rpush(db *DB, args []string) (string, bool) {
	if len(args) < 2 {
		return get_error_string("ERR wrong number of arguments for 'rpush' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	var l *list.List
	if o == nil {
		l = list.New()
		db.dict.Set(key, &Object{
			t: 2,
			p: l,
		})
	} else if o.t != 2 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	} else {
		l = o.p.(*list.List)
	}

	for i := 1; i < len(args); i++ {
		l.PushBack(args[i])
	}

	return get_number_string(l.Len()), true
}

func register_list_cmds() {
	cmds["lpush"] = cmd_lpush
	cmds["lpop"] = cmd_lpop
	cmds["rpush"] = cmd_rpush
	cmds["rpop"] = cmd_rpop
}
