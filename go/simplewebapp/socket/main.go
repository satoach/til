package main

import "flag"

func main() {
	port := ":7777"
	flag.Parse()
	switch flag.Arg(0) {
	case "ip":
		verifyIP(flag.Arg(1))
	case "srv":
		doServ(port)
	case "cli":
		doClient(port, flag.Arg(1))
	default:
		usage()
	}
}
