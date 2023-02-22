package main

import "fmt"

type Student struct {
	name string
	age  int
}

var Info = map[string]bool{
	"001": true,
	"002": true,
}

func main() {
	info := "003"
	s, err := test(info)
	if err != nil {
		return
	}
	student := Student{}
	fmt.Println("student:", student)
	fmt.Printf("%d", 0x80)
	fmt.Println(s)
}
func test(infoo string) (Student, error) {
	if !Info[infoo] {
		fmt.Println("heihei")
		fmt.Println(Student{})
		return Student{}, nil
	}
	return Student{
		name: "test",
	}, nil
}
