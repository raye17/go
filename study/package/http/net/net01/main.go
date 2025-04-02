package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("unKnown")
		panic(err)
	}
	for _, i := range interfaces {
		if len(i.HardwareAddr) > 0 {
			fmt.Println(i.HardwareAddr.String())
		}
	}
}
