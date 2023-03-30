package main

import "fmt"

type man interface {
	getName() string
	setName(string)
}
type stu struct {
	name string
}

func (s *stu) getName() string {
	return s.name
}
func (s *stu) setName(name string) {
	s.name = name
}
func main() {
	var i man
	s := &stu{}
	i = s
	i.setName("001")
	fmt.Println(i.getName())
}
