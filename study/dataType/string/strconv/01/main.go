package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := strconv.FormatFloat(1, 'f', -1, 64)
	b := strconv.FormatFloat(10.23, 'f', -1, 64)
	c := strconv.FormatFloat(10.234567, 'f', -1, 64)
	d := strconv.FormatFloat(0, 'f', -1, 64)

	e := strconv.FormatFloat(1, 'f', 2, 64)
	f := strconv.FormatFloat(1.00, 'f', 2, 64)
	g := strconv.FormatFloat(10.234567, 'f', 2, 64)
	h := strconv.FormatFloat(10.23, 'f', 2, 64)
	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println(strings.TrimSuffix(a, ".00"))
	fmt.Println(strings.TrimSuffix(b, ".00"))
	fmt.Println(strings.TrimSuffix(c, ".00"))
	fmt.Println(strings.TrimSuffix(d, ".00"))
	fmt.Println(strings.TrimSuffix(e, ".00"))
	fmt.Println(strings.TrimSuffix(f, ".00"))
	fmt.Println(strings.TrimSuffix(g, ".00"))
	fmt.Println(strings.TrimSuffix(h, ".00"))
}
