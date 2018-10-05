package main

import (
	"io/ioutil"
	"time"
)

type AOF struct {
	filename string
	cache    []string
	c        chan *Request
	mode     string
}

var g_aof *AOF = &AOF{
	filename: "./aof.log",
	mode:     "AOF_FSYNC_NO",
}

func (aof *AOF) init(mode string, filename string) {
	aof.mode = mode
	aof.filename = filename
	aof.c = make(chan *Request, 1000)
	aof.cache = []string{}
}

func (aof *AOF) write(request *Request) {
	if aof.mode != "AOF_FSYNC_NO" {
		aof.c <- request
	}
}

func (aof *AOF) run() {

	if aof.mode == "AOF_FSYNC_NO" {

	} else if aof.mode == "AOF_FSYNC_ALWAYS" {
		aof.always()
	} else if aof.mode == "AOF_FSYNC_EVERYSEC" {
		aof.everysec()
	}
}

func (aof *AOF) always() {
	for {
		select {
		case r := <-aof.c:
			str := get_request_string(r.cmd, r.args)
			err := ioutil.WriteFile(aof.filename, []byte(str), 0644)
			if err != nil {

			}
		}
	}
}

func (aof *AOF) everysec() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case r := <-aof.c:
			aof.cache = append(aof.cache, get_request_string(r.cmd, r.args))
		case <-ticker.C:
			for i := 0; i < len(aof.cache); i++ {
				err := ioutil.WriteFile(aof.filename, []byte(aof.cache[i]), 0644)
				if err != nil {

				}
			}
		}
	}
}
