package main

//reflect demo
import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v, v.Name(), v.Kind())
	fmt.Printf("type:%v\n", v)
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
