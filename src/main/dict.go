package main

type Dict struct {
	data map[string]*Object
}

func (dict *Dict) Get(key string) *Object {
	o, ok := dict.data[key]
	if ok {
		return o
	}
	return nil
}

func (dict *Dict) Set(key string, o *Object) {
	dict.data[key] = o
}
