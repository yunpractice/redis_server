package main

import (
	"bufio"
	"net"
)

type Client struct {
	conn   net.Conn
	reader Reader
	writer Writer
}

func (c *Client) init(conn net.Conn) {
	c.conn = conn
	c.reader = Reader{
		conn: conn,
		rd:   bufio.NewReader(conn),
	}
	c.writer = Writer{
		conn: conn,
	}
}

func (c *Client) run() {
	// 读写协程
	c.reader.run()
	c.writer.run()
}
