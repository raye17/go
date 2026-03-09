package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	//n := 64
	now := time.Now()
	sum := big.NewInt(1)
	sum = sum.Lsh(sum, 1000000)
	fmt.Println(sum)
	fmt.Println(time.Now().Sub(now))
}
