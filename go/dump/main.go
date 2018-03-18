package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filepath := "file.txt"
	rawdata, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Panicf("%s read failed", filepath)
	}

	fmt.Println(hex.Dump(rawdata))
	const readsize = 16

	for pos := 0; pos < len(rawdata); pos += readsize {
		s := rawdata[pos : pos+readsize]
		fmt.Println(hex.Dump(s))
	}
}
