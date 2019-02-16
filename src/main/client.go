package main

import (
	"bufio"
    "errors"
	"io"
	"net"
	"strconv"
    "strings"
)

type Client struct {
	conn   net.Conn
	rd   *bufio.Reader
    ch chan string
    db int
}

func (c *Client) init(conn net.Conn) {
	c.conn = conn
	c.rd = bufio.NewReader(conn)
    c.ch = make(chan string)
    c.db = 0
}

func (c *Client) run() {
	for {
		request := c.readRequest()
		if request == nil {
			break
		}

		if request.cmd == "COMMAND" {
			c.conn.Write([]byte(get_status_string("OK")))
			continue
		}

		if request.cmd == "select" {
            if len(request.args) != 1 {
			    c.conn.Write([]byte(get_error_string("wrong number of args")))
                continue
            }
            index := request.args[0]
            i,err := strconv.Atoi(index)
            if err != nil {
			    c.conn.Write([]byte(get_error_string("wrong db index")))
			    continue
            }
            c.db = i
			c.conn.Write([]byte(get_status_string("OK")))
            continue
        }

        dbmgr.cmd_chan <- request

        result := <- c.ch

		c.conn.Write([]byte(result))
	}
	c.conn.Close()
}

func (c *Client) readRequest() *Request {
	line, _, err := c.rd.ReadLine()
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

	cmd, err1 := c.readRequestLine()
	if err1 != nil {
		return nil
	}

    // 转换为小写字母
    cmd = strings.ToLower(cmd)

	request := &Request{
		cmd:  cmd,
		args: []string{},
	    ch: c.ch,
	    db: c.db,
	}

	for i := 0; i < n-1; i++ {
		arg, err2 := c.readRequestLine()
		if err2 != nil {
			return nil
		}
		request.args = append(request.args, arg)
	}
	return request
}

func (c *Client) readRequestLine() (string, error) {
	line, _, err := c.rd.ReadLine()
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
	_, err = io.ReadFull(c.rd, b)
	if err != nil {
		return "", errors.New("wrong proto")
	}
	return string(b[:n]), nil
}
