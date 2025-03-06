package reflectss

import (
	"log"
	"reflect"
	"strings"
)

type Person struct{}

func (p *Person) SayHello()                    {}
func (p *Person) Say(msg string)               {}
func (p *Person) SayBye(msg string, times int) {}
func (p *Person) SayNothing()

// 反射获取方法
func Reflects() {
	var p Person
	typ := reflect.TypeOf(&p)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}
