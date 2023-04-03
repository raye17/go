package operator

import "fmt"

type Operator interface {
	SetA(int)
	SetB(int)
	Result()
}
type operatorFactory interface {
	Create() Operator
}
type operatorBase struct {
	a, b int
}

func (o *operatorBase) SetA(a int) {
	o.a = a
}
func (o *operatorBase) SetB(b int) {
	o.b = b
}

type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &plusOperator{
		operatorBase: &operatorBase{},
	}
}

type plusOperator struct {
	*operatorBase
}

func (o plusOperator) Result() {
	fmt.Println(o.a + o.b)
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &minusOperator{
		operatorBase: &operatorBase{},
	}
}

type minusOperator struct {
	*operatorBase
}

func (o minusOperator) Result() {
	fmt.Println(o.a - o.b)
}

type mulOperator struct {
	*operatorBase
}

func (o mulOperator) Result() {
	fmt.Println(o.a * o.b)
}

type MulOperatorFactory struct {
}

func (MulOperatorFactory) Create() Operator {
	return mulOperator{
		operatorBase: &operatorBase{},
	}
}
