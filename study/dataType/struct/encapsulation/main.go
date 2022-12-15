package main

import (
	"encapsulation/model"
	"fmt"
)

func main() {
	p := model.NewPerson("raye")
	p.Set(18, 9999)
	fmt.Println(p)
	fmt.Println(p.Get())
}
