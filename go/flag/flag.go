package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type intSliceValue []int

func (isv *intSliceValue) Set(str string) error {
	for _, s := range strings.Split(str, ",") {
		v, err := strconv.ParseInt(s, 0, 32)
		if err != nil {
			return err
		}
		*isv = append(*isv, int(v))
	}
	return nil
}

func (isv *intSliceValue) String() string {
	return fmt.Sprintf("%v", *isv)
}

func main() {
	var isTrue bool
	var numbers intSliceValue

	flag.BoolVar(&isTrue, "b", false, "flag test")
	flag.Var(&numbers, "number", "number array test")
	flag.Parse()

	fmt.Println("isTrue:", isTrue)
	fmt.Println(numbers)
}
