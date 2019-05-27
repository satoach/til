package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func doServ(service string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		} else {
			fmt.Printf("accept %v\n", conn)
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	request := make([]byte, 128)
	defer conn.Close()
	for {
		rlen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if rlen == 0 {
			fmt.Printf("no len")
			break
		} else if string(request) == "timestamp" {
			fmt.Printf("timestamp")
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			fmt.Printf("default: %s", string(request))
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // cler last read content
	}
}
