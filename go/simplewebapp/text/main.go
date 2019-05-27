package main

import (
	"flag"
	"fmt"
	"os"
)

type MyText interface {
	Write()
	Read()
}

func usage() {
	fmt.Printf("USASE:\n%v <xml|json> <write|read>", os.Args[0])
}

func getText(name string) (MyText, error) {
	switch name {
	case "xml":
		return &MyXML{}, nil
	case "json":
		return &MyJSON{}, nil
	}
	return nil, fmt.Errorf("'%v' is unsupport", name)
}

func doText(txt MyText, cmd string) error {
	switch cmd {
	case "write":
		txt.Write()
	case "read":
		txt.Read()
	default:
		return fmt.Errorf("'%v' is unsupport", cmd)
	}
	return nil
}

func main() {
	flag.Parse()
	txt, err := getText(flag.Arg(0))
	if err != nil {
		fmt.Printf("%v\n", err)
		usage()
		return
	}

	err = doText(txt, flag.Arg(1))
	if err != nil {
		fmt.Printf("%v\n", err)
		usage()
		return
	}
}
