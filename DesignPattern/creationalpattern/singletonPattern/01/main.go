package main

import (
	"fmt"
	"singletonPattern01/model"
)

func main() {
	m := model.GetStuMsg()
	m.SetMsg("001", 36001, 16)
	a, b, c := m.GetMsg()
	fmt.Println(a, b, c)
	l := model.GetStuMsg()
	x, y, z := l.GetMsg()
	fmt.Println(x, y, z)
}
