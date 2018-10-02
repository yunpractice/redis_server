package main

type DB struct {
	index int   // db序号0-15
	dict  *Dict // hash表
}

type DBMgr struct {
	dbs []*DB
}

var dbmgr *DBMgr = &DBMgr{
	dbs: []*DB{},
}

func (mgr *DBMgr) init(dbNum int) {
	for i := 0; i < dbNum; i++ {
		db := &DB{
			index: i,
			dict:  &Dict{},
		}
		db.dict.init()
		mgr.dbs = append(mgr.dbs, db)
	}
}

func (mgr *DBMgr) getDB(index int) *DB {
	if index < 0 || index >= len(mgr.dbs) {
		return nil
	}
	return mgr.dbs[index]
}
