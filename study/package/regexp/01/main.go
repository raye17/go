package main

import (
	"fmt"
	"regexp"
)

var (
	controlChars   = regexp.MustCompile("[[:cntrl:]]")
	controlEncoded = regexp.MustCompile("%[0-1][0-9,a-f,A-F]")
)

func main() {
	fmt.Println(controlEncoded.FindStringIndex("fgdh%0chdhd%1F"))
	fmt.Printf("q:%q", `%20`)
}
