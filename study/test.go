package main

import (
	"fmt"
	"github.com/golang/glog"
)

func main() {
	fmt.Println("start...")
	i := 1
add:
	if i < 5 {
		glog.Errorf("i<5 i:%d", i)
		i++
		glog.Infof("i:%d", i)
		fmt.Println(i)
		goto add
	}
}
