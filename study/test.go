package main

import (
	"fmt"
)

type i interface {
	get()
	set(string, int)
}
type stu struct {
	name string
	age  int
}

func (s stu) get() {
	fmt.Println(s.name, ":", s.age)
}
func (s stu) set(name string, age int) {
	s.name = name
	s.age = age
}

type teacher struct {
	name string
	age  int
}

func (t *teacher) get() {
	fmt.Println(t.name, t.age)
}
func (t *teacher) set(name string, age int) {
	t.name = name
	t.age = age
}

type inter01 interface {
	get() string
}
type str01 struct {
	inter01
	name string
}

func (s *str01) get() string {
	ss := s.inter01.get()
	fmt.Println(ss)
	return ss
}
func main() {
	//fmt.Println(mul(7) + mul(6) + mul(5) + mul(4) + mul(3) + mul(2) + mul(1) + mul(0))
	//fmt.Println(mul(31))
	//sli01 := make([]string, 7)
	//sli01 = []string{"Go", "Before", "Linux", "Python", "java", "after"}
	//sort.Strings(sli01)
	//fmt.Println(sli01)
	//var cou int64
	//d := atomic.AddInt64(&cou, 1)
	//fmt.Println(d)
	//fmt.Printf("%c", 125)
	//d1 := 4
	////d2 := (d1 >> 1) & 3
	//fmt.Println(d1 & 3)
	s := []int{96, 90, 41, 82, 39, 74, 64, 50, 30}
	maxScore(s, 8)
}
func maxScore(cardPoints []int, k int) int {
	count, result, length := 0, 0, len(cardPoints)
	for _, v := range cardPoints[length-k:] {
		count += v
	}
	result = count
	for i := 0; i < k; i++ {
		count = count + cardPoints[i] - cardPoints[length-k+i]
		if count > result {
			result = count
		}
	}
	fmt.Println(result)
	return result
}
func mul(i int) int {
	sum := 1
	for j := 0; j < i; j++ {
		sum = sum * 2
	}
	return sum
}
