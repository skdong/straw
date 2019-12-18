package main

import (
	"fmt"
	"net"
)

func callServer() error {
	conn, err := net.Dial("tcp", ":8080")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	conn.Write([]byte("hello world"))
	return nil
}

func main() {
	callServer()
}
