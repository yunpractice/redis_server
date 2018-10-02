package main

func cmd_get(db *DB, args []string) string {
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {

	}
	if o.t {

	}
}
