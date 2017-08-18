package main

import (
	"flag"
	"fmt"
)

func main() {
	var isTrue bool
	flag.BoolVar(&isTrue, "v", false, "flag test")
	flag.Parse()
	fmt.Println("isTrue:", isTrue)
	return
}
