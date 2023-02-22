package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "raye:sxy:lcx:i love novel,welcome to my host"
	s1 := strings.SplitN(s, ":", -1)
	index := strings.Index(s, "y")
	lastindex := strings.LastIndex(s, "sxr")
	fmt.Println(s1, index, lastindex)
}
