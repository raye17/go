package main

import "fmt"

type Brand int

const (
	HuaWei Brand = iota
	Apple
	Xm
	Vivo
)

type Phone interface {
	ShowBrand()
}
type Iphone struct{}

func (i *Iphone) ShowBrand() {
	fmt.Println("apple...")
}

type HPhone struct{}

func (h *HPhone) ShowBrand() {
	fmt.Println("华为...")
}

type XPhone struct{}

func (x *XPhone) ShowBrand() {
	fmt.Println("小米...")
}

type VPhone struct{}

func (v *VPhone) ShowBrand() {
	fmt.Println("vivo...")
}
func NewPhone(brand Brand) Phone {
	switch brand {
	case HuaWei:
		return &HPhone{}
	case Apple:
		return &Iphone{}
	case Xm:
		return &XPhone{}
	case Vivo:
		return &VPhone{}
	default:
		return nil
	}
}
func main() {
	var phone Phone
	//华为
	phone = NewPhone(HuaWei)
	phone.ShowBrand()
	//vivo
	phone = NewPhone(Vivo)
	phone.ShowBrand()
}
