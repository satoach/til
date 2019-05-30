package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func doClient(port string) {
	client, err := rpc.DialHTTP("tcp", ":"+port)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("arith:", err)
	}
	fmt.Printf("Arith: %d/%d=%d, remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)
}
