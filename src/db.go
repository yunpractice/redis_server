package main

type DB struct {
	index int   // db序号0-15
	dict  *Dict // hash表
}
