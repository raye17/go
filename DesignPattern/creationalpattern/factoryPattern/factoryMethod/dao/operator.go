package dao

import "factoryMethod/model/operator"

func Plus(a, b int) {
	p := operator.PlusOperatorFactory{}.Create()
	p.SetA(a)
	p.SetB(b)
	p.Result()
}
func Min(a, b int) {
	m := operator.MinusOperatorFactory{}.Create()
	m.SetA(a)
	m.SetB(b)
	m.Result()
}
func Mul(a, b int) {
	m := operator.MulOperatorFactory{}.Create()
	m.SetA(a)
	m.SetB(b)
	m.Result()
}
