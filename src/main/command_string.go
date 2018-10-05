package main

import (
	"strconv"
)

func cmd_append(db *DB, args []string) (string, bool) {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'append' command"), false
	}

	key := args[0]
	value := args[1]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: value,
		})
		return get_number_string(len(value)), true
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	newstr := o.p.(string) + value
	o.p = newstr
	return get_number_string(len(newstr)), true
}

func cmd_decr(db *DB, args []string) (string, bool) {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'decr' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: "-1",
		})
		return get_number_string(-1), true
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	n, err := strconv.Atoi(o.p.(string))
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range"), false
	}
	n--
	o.p = string(n)
	return get_number_string(n), true
}

func cmd_decrby(db *DB, args []string) (string, bool) {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'decrby' command"), false
	}
	key := args[0]
	by, err := strconv.Atoi(args[1])
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range"), false
	}
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: string(-by),
		})
		return get_number_string(-by), true
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	n, err1 := strconv.Atoi(o.p.(string))
	if err1 != nil {
		return get_error_string("ERR value is not an integer or out of range"), false
	}
	n = n - by
	o.p = string(n)
	return get_number_string(n), true
}

func cmd_get(db *DB, args []string) (string, bool) {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'get' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_empty_bulk_string(), false
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	return get_bulk_string(o.p.(string)), false
}

func cmd_getset(db *DB, args []string) (string, bool) {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'getset' command"), false
	}
	key := args[0]
	value := args[1]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: value,
		})
		return get_empty_bulk_string(), true
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	o.p = value
	return get_bulk_string(value), true
}

func cmd_incr(db *DB, args []string) (string, bool) {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'incr' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: "1",
		})
		return get_number_string(1), true
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	n, err := strconv.Atoi(o.p.(string))
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range"), false
	}
	n++
	o.p = string(n)
	return get_number_string(n), true
}

func cmd_incrby(db *DB, args []string) (string, bool) {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'incrby' command"), false
	}
	key := args[0]
	by, err := strconv.Atoi(args[1])
	if err != nil {
		return get_error_string("ERR value is not an integer or out of range"), false
	}
	o := db.dict.Get(key)
	if o == nil {
		db.dict.Set(key, &Object{
			t: 1,
			p: string(by),
		})
		return get_number_string(by), true
	}
	if o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}
	n, err1 := strconv.Atoi(o.p.(string))
	if err1 != nil {
		return get_error_string("ERR value is not an integer or out of range"), false
	}
	n = n + by
	o.p = string(n)
	return get_number_string(n), true
}

func cmd_mget(db *DB, args []string) (string, bool) {
	if len(args) < 1 {
		return get_error_string("ERR wrong number of arguments for 'mget' command"), false
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
	return get_bulks_string(result), false
}

func cmd_set(db *DB, args []string) (string, bool) {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'set' command"), false
	}
	key := args[0]
	value := args[1]
	db.dict.Set(key, &Object{
		t: 1,
		p: value,
	})
	return get_status_string("OK"), true
}

func cmd_setnx(db *DB, args []string) (string, bool) {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'setnx' command"), false
	}
	key := args[0]
	value := args[1]
	if db.dict.Get(key) != nil {
		return get_number_string(0), false
	}
	db.dict.Set(key, &Object{
		t: 1,
		p: value,
	})
	return get_number_string(1), true
}

func cmd_strlen(db *DB, args []string) (string, bool) {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'strlen' command"), false
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil || o.t != 1 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value"), false
	}

	return get_number_string(len(o.p.(string))), false
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
