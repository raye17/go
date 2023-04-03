package main

import "fmt"

func main() {
	Pay()
}
func Pay() {
	fmt.Println("please choose your pay style:")
	fmt.Println(`
	1:alipay
	2:vx
	`)
	var n int
	fmt.Scanln(&n)
	switch n {
	case 1:
		pay := NewPayPlatform("alipay")
		pay.Pay()
	case 2:
		pay := NewPayPlatform("vx")
		pay.Pay()
	}
}

type PayInterface interface {
	Pay()
}
type AliPay struct {
}

func (*AliPay) Pay() {
	fmt.Println("Alipay...")
}

type VX struct {
}

func (*VX) Pay() {
	fmt.Println("vx...")
}
func NewPayPlatform(platform string) PayInterface {
	switch platform {
	case "alipay":
		return &AliPay{}
	case "vx":
		return &VX{}
	default:
		return &AliPay{}
	}
}
