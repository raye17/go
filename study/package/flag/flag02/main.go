package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	var p string
	s := flag.String("ben", "node", "usage-ben")
	flag.StringVar(&p, "pp", "v-pp", "usage-pp")
	flag.Parse()
	fmt.Println(*s)
	fmt.Println(p)

}
