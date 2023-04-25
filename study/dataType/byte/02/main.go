package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	ipB := [4]byte{192, 168, 51, 90}
	p1 := math.MaxInt8
	p2 := math.MaxUint8
	fmt.Println(p1, p2)
	var ip string
	var iip []string
	for _, v := range ipB {
		v1 := strconv.Itoa(int(v))
		iip = append(iip, v1)
	}
	ip = strings.Join(iip, ".")
	fmt.Println(ip)
}
