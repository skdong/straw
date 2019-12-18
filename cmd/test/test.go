package main

import (
	"errors"
	"fmt"
	"net"
	"path"
	"time"
)

func sleep() {
	time.Sleep(10000 * time.Millisecond)
}

func scan(ip, port string) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), 60)
	if err != nil {
		fmt.Println(err)
	} else {
		defer conn.Close()
	}
}

func splithostport() {
	host, port, err := net.SplitHostPort("localhost:8080")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(host, port)
	}
}

func parsecdir() {
	ip, net, err := net.ParseCIDR("10.1.0.0/24")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ip, net)
	}
}

func parseip() {
	ip := net.ParseIP("10.1.0.1")
	fmt.Println(ip)
}

func bytestring() {
	buffer := []byte("hello world")
	fmt.Println(buffer)
	fmt.Println(string(buffer))
}

func absolutePath() {
	pwd := "C://test"
	fmt.Println(path.Join(pwd, "aa"))
}

func newError() {
	fmt.Println(errors.New("aa"))
}

func goRutine() {
	ch := make(chan int)
	defer close(ch)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("test")
			ch <- i
		}()
	}
	var ret []int
	for len(ret) != 100 {
		ret = append(ret, <-ch)
	}
}

func main() {
	goRutine()
}
