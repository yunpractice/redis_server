package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
)

type Reader struct {
	conn net.Conn
	rd   *bufio.Reader
}

func (r *Reader) run() {
	for {
		request := r.readRequest()
		if request == nil {
			break
		}
		fmt.Println(request)

		if request.cmd == "COMMAND" {
			buf := "+OK\r\n"
			r.conn.Write([]byte(buf))
			continue
		}

		db := dbmgr.getDB(0)
		result, needaof := dispatch_cmd(db, request.cmd, request.args)
		if needaof {
			g_aof.write(request)
		}
		r.conn.Write([]byte(result))
		fmt.Println(result)
	}
	r.conn.Close()
}

func (r *Reader) readRequest() *Request {
	line, _, err := r.rd.ReadLine()
	if err != nil {
		return nil
	}

	if line[0] != '*' {
		return nil
	}

	// 命令行数
	n := 0
	n, err = strconv.Atoi(string(line[1:]))
	if n <= 0 || n > 1000 {
		return nil
	}

	cmd, err1 := r.readRequestLine()
	if err1 != nil {
		return nil
	}

	request := &Request{
		cmd:  string(cmd),
		args: []string{},
	}

	for i := 0; i < n-1; i++ {
		arg, err2 := r.readRequestLine()
		if err2 != nil {
			return nil
		}
		request.args = append(request.args, arg)
	}
	return request
}

func (r *Reader) readRequestLine() (string, error) {
	line, _, err := r.rd.ReadLine()
	if err != nil {
		return "", err
	}
	if line[0] != '$' {
		return "", errors.New("wrong proto")
	}
	n := 0
	n, err = strconv.Atoi(string(line[1:]))
	if err != nil {
		return "", err
	}
	if n <= 0 {
		return "", errors.New("wrong proto")
	}
	b := make([]byte, n+2)
	_, err = io.ReadFull(r.rd, b)
	if err != nil {
		return "", errors.New("wrong proto")
	}
	return string(b[:n]), nil
}
