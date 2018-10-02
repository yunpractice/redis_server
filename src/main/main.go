package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("need config file name!")
	}

	file := os.Args[1]
	fmt.Println("config file is " + file)
	config.Load(file)
}
