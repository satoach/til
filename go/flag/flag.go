package main

import (
	"flag"
	"fmt"
	"os"
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

func subcmd() {
	sub := flag.NewFlagSet(os.Args[0]+" "+os.Args[1], flag.ExitOnError)
	var isTrue bool

	sub.BoolVar(&isTrue, "b", false, "flag test")
	sub.Parse(os.Args[2:])

	fmt.Println("sub: isTrue:", isTrue)
}

func maincmd() {
	var isTrue bool
	var numbers intSliceValue

	flag.BoolVar(&isTrue, "b", false, "flag test")
	flag.Var(&numbers, "number", "number array test")
	flag.Parse()

	fmt.Println("isTrue:", isTrue)
	fmt.Println(numbers)
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "sub":
			subcmd()
			return
		default:
			break
		}
	}
	maincmd()
	return
}
