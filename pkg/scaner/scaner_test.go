package scaner

import (
	"testing"
	"net"
	"github.com/stretchr/testify/assert"
)

var runFlag bool = true

func runServer(service string){
	ln, err := net.Listen("tcp", service)
	if err != nil{
		return
	}
	for runFlag == true{
		conn, _ := ln.Accept()
		conn.Close()
	}
}


func TestScan(t * testing.T){
	host, port := "localhost", "8083"
	go runServer(net.JoinHostPort(host, port))	
	assert.Equal(t, nil, scan(host, port))
	runFlag = false
}

func TestScanPorts(t * testing.T){
	host := "localhost"
	ports := []string{"8080", "80"}
	scanPorts(host, ports)
}

func TestScanHostsPorts(t * testing.T){
	hosts := []string{"localhost", "10.0.1.10"}
	ports := []string{"8080", "80"}
	scanHostsPorts(hosts, ports)
}