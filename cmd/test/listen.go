package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn){
	defer conn.Close()
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(len)
	fmt.Println(string(buffer))
}

func startServer() error {
	ln, err := net.Listen("tcp", ":8080")
	defer ln.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for{
		conn, err := ln.Accept()
		if err!= nil{
			fmt.Println(err)
			return err
		}
		go handleConnection(conn)
	}
}

func main(){
	if err := startServer() ; err != nil{
		fmt.Println(err)
	}
}