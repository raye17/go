package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var portMap = make(map[int]int)
	portMap[20001] = 80
	pb, _ := json.Marshal(portMap)
	fmt.Println(pb)
	fmt.Println(string(pb))
	sp := string(pb)
	fmt.Println("**********")
	var p map[int]int
	err := json.Unmarshal([]byte(sp), &p)
	if err != nil {
		fmt.Println(p)
		return
	}
	fmt.Println(p)
}
