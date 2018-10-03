package main

import (
	"strconv"
)

func cmd_append(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'append' command")
	}

	key := args[0]
	value := args[1]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: value,
		})
		return get_number_string(len(value))
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	newstr := o.p.(string) + value
	o.p = newstr
	return get_number_string(len(newstr))
}

func cmd_decr(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'decr' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: "-1",
		})
		return get_number_string(-1)
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	n, err := strconv.Atoi(o.p.(string))
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range")
	}
	n--
	o.p = string(n)
	return get_number_string(n)
}

func cmd_decrby(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'decrby' command")
	}
	key := args[0]
	by, err := strconv.Atoi(args[1])
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range")
	}
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: string(-by),
		})
		return get_number_string(-by)
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	n, err1 := strconv.Atoi(o.p.(string))
	if err1 != nil {
		return get_error_string("ERR value is not an integer or out of range")
	}
	n = n - by
	o.p = string(n)
	return get_number_string(n)
}

func cmd_get(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'get' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_empty_bulk_string()
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	return get_bulk_string(o.p.(string))
}

func cmd_getset(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'getset' command")
	}
	key := args[0]
	value := args[1]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: value,
		})
		return get_empty_bulk_string()
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	o.p = value
	return get_bulk_string(value)
}

func cmd_incr(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'incr' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: "1",
		})
		return get_number_string(1)
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	n, err := strconv.Atoi(o.p.(string))
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range")
	}
	n++
	o.p = string(n)
	return get_number_string(n)
}

func cmd_incrby(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'incrby' command")
	}
	key := args[0]
	by, err := strconv.Atoi(args[1])
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range")
	}
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: string(by),
		})
		return get_number_string(by)
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	n, err1 := strconv.Atoi(o.p.(string))
	if err1 != nil {
		return get_error_string("ERR value is not an integer or out of range")
	}
	n = n + by
	o.p = string(n)
	return get_number_string(n)
}

func cmd_mget(db *DB, args []string) string {
	if len(args) < 1 {
		return get_error_string("ERR wrong number of arguments for 'mget' command")
	}

	result := []string{}
	for i := 0; i < len(args); i++ {
		key := args[i]
		o := db.dict.Get(key)
		if o == nil || o.t != 1 {
			result = append(result, get_empty_bulk_string())
		} else {
			result = append(result, get_bulk_string(o.p.(string)))
		}

	}
	return get_bulks_string(result)
}

func cmd_set(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'set' command")
	}
	key := args[0]
	value := args[1]
	db.dict.Set(key, &Object{
		t: 1,
		p: value,
	})
	return get_status_string("OK")
}

func cmd_setnx(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'setnx' command")
	}
	key := args[0]
	value := args[1]
	if db.dict.Get(key) != nil {
		return get_number_string(0)
	}
	db.dict.Set(key, &Object{
		t: 1,
		p: value,
	})
	return get_number_string(1)
}

func cmd_strlen(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'strlen' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil || o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	return get_number_string(len(o.p.(string)))
}

func register_string_cmds() {
	cmds["append"] = cmd_append
	cmds["decr"] = cmd_decr
	cmds["decrby"] = cmd_decrby
	cmds["get"] = cmd_get
	cmds["getset"] = cmd_getset
	cmds["incr"] = cmd_incr
	cmds["incrby"] = cmd_incrby
	cmds["mget"] = cmd_mget
	cmds["set"] = cmd_set
	cmds["setnx"] = cmd_setnx
	cmds["strlen"] = cmd_strlen
}
