package main

import "fmt"

func main() {
	adaptee := &adapteeImpl{}
	target := NewAdapter(adaptee)
	fmt.Println(target.Request())
}

type Target interface {
	Request() string
}
type Adaptee interface {
	SpecificQuest() string
}
type adapteeImpl struct {
}

func (a *adapteeImpl) SpecificQuest() string {
	return "Adaptee method"
}

type adapter struct {
	adaptee Adaptee
}

func (a *adapter) Request() string {
	return fmt.Sprintf("Adapter method :%s", a.adaptee.SpecificQuest())
}
func NewAdapter(adaptee Adaptee) Target { return &adapter{adaptee: adaptee} }
