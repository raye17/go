package main

import "fmt"

type vx struct {
}
type alipay struct {
}
type pay interface {
	pay()
}

func (vx vx) pay() {
	fmt.Println("vx pay...")
}
func (alipay alipay) pay() {
	fmt.Println("alipay pay...")
}
func choosePay(p pay) {
	//fmt.Println("choose ")
	p.pay()
}
func main() {
	vx1 := vx{}
	choosePay(vx1)
}
