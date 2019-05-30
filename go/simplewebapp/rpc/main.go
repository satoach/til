package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("USAGE:\n%v <srv|cli>\n", os.Args[0])
}

func main() {
	port := "7777"
	flag.Parse()
	switch flag.Arg(0) {
	case "srv":
		doServ(port)
	case "cli":
		doClient(port)
	default:
		usage()
	}
}
