package main

func cmd_sadd(db *DB, args []string) string {
	if len(args) < 2 {
		return get_error_string("ERR wrong number of arguments for 'rpush' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	var s map[string]bool
	if o == nil {
		s = map[string]bool{}
		db.dict.Set(key, &Object{
			t: 3,
			p: s,
		})
	} else if o.t != 3 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	} else {
		s = o.p.(map[string]bool)
	}

	n := 0
	for i := 1; i < len(args); i++ {
		_, ok := s[args[i]]
		if !ok {
			s[args[i]] = true
			n++
		}
	}

	return get_number_string(n)
}

func cmd_scard(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'scard' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_number_string(0)
	} else if o.t != 3 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	s := o.p.(map[string]bool)

	return get_number_string(len(s))
}

func cmd_sdiff(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'sdiff' command")
	}
	key1 := args[0]
	key2 := args[1]
	o1 := db.dict.Get(key1)
	o2 := db.dict.Get(key2)

	var s map[string]bool

	if o1 == nil && o2 == nil {
		return get_empty_bulks_string()
	} else if o1 != nil && o1.t != 3 || o2 != nil && o2.t != 3 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	} else if o1 == nil {
		s = o2.p.(map[string]bool)
	} else if o2 == nil {
		s = o1.p.(map[string]bool)
	}

	result := []string{}
	if s != nil {
		for k := range s {
			result = append(result, k)
		}
	} else {
		s1 := o1.p.(map[string]bool)
		s2 := o2.p.(map[string]bool)
		for k := range s1 {
			_, ok := s2[k]
			if !ok {
				result = append(result, k)
			}
		}
		for k := range s2 {
			_, ok := s1[k]
			if !ok {
				result = append(result, k)
			}
		}
	}

	return get_bulks_string(result)
}

func cmd_sismember(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'sismember' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	var s map[string]bool
	if o == nil {
		return get_number_string(0)
	} else if o.t != 3 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	s = o.p.(map[string]bool)

	_, ok := s[args[1]]
	if ok {
		return get_number_string(1)
	}

	return get_number_string(0)
}

func cmd_smembers(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("ERR wrong number of arguments for 'smembers' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	var s map[string]bool
	if o == nil {
		return get_number_string(0)
	} else if o.t != 3 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	s = o.p.(map[string]bool)

	result := []string{}
	for k := range s {
		result = append(result, k)
	}

	return get_bulks_string(result)
}

func cmd_spop(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("ERR wrong number of arguments for 'spop' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	var s map[string]bool
	if o == nil {
		return get_number_string(0)
	} else if o.t != 3 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	s = o.p.(map[string]bool)

	// to-do
	_ = s

	return get_number_string(0)
}

func cmd_srem(db *DB, args []string) string {
	if len(args) < 2 {
		return get_error_string("ERR wrong number of arguments for 'srem' command")
	}
	key := args[0]
	o := db.dict.Get(key)
	var s map[string]bool
	if o == nil {
		return get_number_string(0)
	} else if o.t != 3 {
		return get_error_string("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	s = o.p.(map[string]bool)

	n := 0
	for i := 1; i < len(args); i++ {
		_, ok := s[args[i]]
		if ok {
			delete(s, args[i])
			n++
		}
	}

	return get_number_string(n)
}

func register_set_cmds() {
	cmds["sadd"] = cmd_sadd
	cmds["scard"] = cmd_scard
	cmds["sdiff"] = cmd_sdiff
	cmds["smembers"] = cmd_smembers
	cmds["spop"] = cmd_scard
	cmds["srem"] = cmd_scard
}
