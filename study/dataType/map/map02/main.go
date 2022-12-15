package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	//map初始化
	var a = make(map[string]int, 8)
	fmt.Println(a == nil)
	a["raye"] = 17
	a["lcx"] = 18
	a["sxy"] = 21
	a["xxx"] = 3
	fmt.Println(a)
	fmt.Printf("a:%#v\n", a)
	value, ok := a["raye"]
	if ok {
		fmt.Println("sxy在", value)
	} else {
		fmt.Println("wu")
	}
	//map的遍历
	for k, v := range a {
		fmt.Println(k, v)
	}
	delete(a, "xxx")
	fmt.Println(a)
	var b = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("sud%02d", i)
		value := rand.Intn(100)
		b[key] = value
	}
	// for k, v := range b {
	// 	fmt.Println(k, v)
	// }
	keys := make([]string, 0, 100)
	for k := range b {
		keys = append(keys, k)
		sort.Strings(keys)
	}
	for _, key := range keys {
		fmt.Println(key, b[key])
	}
}
