package main

import (
	"container/list"
)

func cmd_lpop(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'lpop' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_empty_bulk_string()
	} else if o.t != 2 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	l := o.p.(*list.List)
	if l.Len() == 0 {
		return get_empty_bulk_string()
	}

	node := l.Front()
	l.Remove(node)

	return get_bulk_string(node.Value.(string))
}

func cmd_lpush(db *DB, args []string) string {
	if len(args) < 2 {
		return get_error_string("ERR wrong number of arguments for 'lpush' command")
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
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	} else {
		l = o.p.(*list.List)
	}

	for i := 1; i < len(args); i++ {
		l.PushFront(args[i])
	}

	return get_number_string(l.Len())
}

func cmd_rpop(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'rpop' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_empty_bulk_string()
	} else if o.t != 2 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	l := o.p.(*list.List)
	if l.Len() == 0 {
		return get_empty_bulk_string()
	}

	node := l.Back()
	l.Remove(node)

	return get_bulk_string(node.Value.(string))
}

func cmd_rpush(db *DB, args []string) string {
	if len(args) < 2 {
		return get_error_string("ERR wrong number of arguments for 'rpush' command")
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
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	} else {
		l = o.p.(*list.List)
	}

	for i := 1; i < len(args); i++ {
		l.PushBack(args[i])
	}

	return get_number_string(l.Len())
}

func register_list_cmds() {
	cmds["lpush"] = cmd_lpush
	cmds["lpop"] = cmd_lpop
	cmds["rpush"] = cmd_rpush
	cmds["rpop"] = cmd_rpop
}
