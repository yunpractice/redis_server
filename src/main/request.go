package main

type Request struct {
	cmd  string
	args []string
    ch chan string
    db int
}
