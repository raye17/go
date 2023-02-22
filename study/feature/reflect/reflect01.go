package main

//reflect demo
import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t, t.Name(), t.Kind())
	fmt.Printf("type:%v\n", t)
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)

	fmt.Printf("value:%v\n", v)
}
func main() {
	var a float32 = 3.14
	//var b int64 = 100
	reflectType(a)
	//reflectType(b)
	//reflectValue(a)
	//reflectValue(b)

}
