package main

type DB struct {
	index int   // db序号0-15
	dict  *Dict // hash表
}

type DBMgr struct {
	dbs []*DB
    cmd_chan chan *Request
}

var dbmgr *DBMgr = &DBMgr{
	dbs: []*DB{},
    cmd_chan: make(chan *Request,100),
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

func (mgr *DBMgr) run(){
    for {
        select {
        case request := <- mgr.cmd_chan:
            db := mgr.getDB(request.db)
            if db == nil {
                request.ch <- ""
            }else{
                result, needaof := dispatch_cmd(db,request.cmd,request.args)
                if needaof {
			        g_aof.write(request)
		        }
		        request.ch <- result
            }
        }
    }
}

func (mgr* DBMgr) shutdown(){
}
