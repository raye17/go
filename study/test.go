package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	a := resource.MustParse("2.5m")
	fmt.Println(a)
	fmt.Printf("%#v", a)
}
