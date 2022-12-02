package main

import (
	"flag"
	"fmt"
)

var OK bool

func main() {
	flag.BoolVar(&OK, "h", false, "print ok")
	flag.Parse()
	if OK {
		fmt.Println("help")
	} else {
		fmt.Println("Not help")
	}
}
