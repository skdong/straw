package scaner

import (
	"errors"
	"fmt"
	"net"
)

func scan(host, port string) error {
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println(err)
		return errors.New("Connect Error")
	} else {
		defer conn.Close()
		fmt.Println("Connect OK")
	}
	return nil
}

func scanPorts(host string, ports []string) error {
	for i := range ports {
		go scan(host, ports[i])
	}
	return nil
}

func scanHostsPorts(hosts []string, ports []string) error {
	for i := range hosts {
		go scanPorts(hosts[i], ports)
	}
	return nil
}
