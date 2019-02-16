package main

import (
	"fmt"
	"net"
)

type Listener struct {
	l       net.Listener
	id      int
	clients map[int]*Client
}

func (listener *Listener) init(port string) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	listener.l = l
	listener.id = 0
	listener.clients = map[int]*Client{}
}

func (listener *Listener) run() {
	for {
		c, err := listener.l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		listener.handleConn(c)
	}
}

func (listener *Listener) handleConn(conn net.Conn) {
	fmt.Println("new Connection", conn.RemoteAddr())
	client := &Client{}
	listener.id++
	listener.clients[listener.id] = client

    client.init(conn)
	client.run()
}

func (listener *Listener) shutdown(){
    // 不再接收新连接
    listener.l.Close()
}
