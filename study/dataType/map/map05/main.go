package main

import (
	"fmt"
)

type m map[interface{}]interface{}
type mm []m

func main() {
	m1 := make(m)
	m2 := make(m, 2)
	m3 := m{"m3-01": 1, "m3-02": 2}
	fmt.Println(m1, len(m1))
	fmt.Println(m2, len(m2))
	m1["001"] = "m1-01"
	m1["002"] = "m1-02"
	m1["000"] = "m1-00"
	m2[19] = 91
	m2[18] = 81
	m1["003"] = "m1-03"
	var mms mm
	m4 := m{"inter01": "inter-01"}
	m4["in02"] = 222
	mms = append(mms, m4)
	mms = append(mms, m1)
	mms = append(mms, m2)
	mms = append(mms, m3)
	prints(mms)
}
func prints(m mm) {
	for _, v := range m {
		fmt.Println("...", v)
		for l, k := range v {
			fmt.Println(l, k)
		}
	}
}
