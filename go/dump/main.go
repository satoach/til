package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var filepath string
	flag.StringVar(&filepath, "p", "file.txt", "file path")
	flag.Parse()

	rawdata, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("%s read failed", filepath)
	}

	fmt.Println(hex.Dump(rawdata))
	const readsize = 16

	for pos := 0; pos < len(rawdata); pos += readsize {
		s := rawdata[pos : pos+readsize]
		fmt.Println(hex.Dump(s))
	}
}
