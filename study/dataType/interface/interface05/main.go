package main

import "fmt"

// ISay 定义ISay接口，该接口实现了Say方法
type ISay interface {
	Say()
	Run()
}

// SayFunc 定义函数类型的sayFunc
type SayFunc func()

func (s SayFunc) Say() {
	fmt.Println("使用了SayFunc里的Say()")
	s()
}
func (s SayFunc) Run() {
	fmt.Println("run")
}

// 入口函数具体执行的函数
func say(iSay ISay) {
	fmt.Println("使用了say")
	iSay.Say()
	iSay.Run()
}

// Say 入口函数
func Say(handler func()) {
	fmt.Println("将传进来的handler转为SayFunc类型")
	say(SayFunc(handler))
}

// SayHello 定义一个函数
func SayHello() {
	fmt.Println("hello")
}
func main() {
	Say(SayHello)
}
