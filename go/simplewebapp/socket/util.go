package main

import (
	"fmt"
	"net"
	"os"
)

func usage() {
	fmt.Printf("USAGE:\n%v <ip|srv|cli>\n", os.Args[0])
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func verifyIP(name string) {
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid addr")
	} else {
		fmt.Println("addr:", addr.String())
	}

	os.Exit(0)
}
