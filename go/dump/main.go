package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func byte2Str(b []byte, max int) (string, error) {
	end := bytes.Index(b, []byte{0})
	if end >= max {
		err := fmt.Errorf("end[%d] >= max[%d]", end, max)
		return "", err
	}
	s := string(b[:end])
	return s, nil
}

func conv2BigEndian(b []byte) uint32 {
	return binary.BigEndian.Uint32(b[0:4])
}

func main() {
	var filepath string
	flag.StringVar(&filepath, "p", "file.txt", "file path")
	flag.Parse()

	rawdata, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("%s read failed", filepath)
	} else if len(rawdata) < 16 {
		log.Fatalf("%s size: %d", filepath, len(rawdata))
	}

	s, err := byte2Str(rawdata, 8)
	if err != nil {
		log.Fatalf("s: %s", err)
	}
	fmt.Printf("%s: %d", s, conv2BigEndian(rawdata[8:16]))

	fmt.Println(hex.Dump(rawdata))
	const readsize = 16

	for pos := 0; pos < len(rawdata); pos += readsize {
		s := rawdata[pos : pos+readsize]
		fmt.Println(hex.Dump(s))
	}
}
