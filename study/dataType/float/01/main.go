package main

import "fmt"

func main() {
	f := []float32{12.00, 17.18, 13.98, 15.45}
	for _, v := range f {
		fmt.Println(fmt.Sprintf("%.2f", v))
	}
}
