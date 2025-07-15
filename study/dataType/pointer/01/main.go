package main

func main() {
	var a int = 10
	var b *int = &a
	var s string = "hello"
	var s1 *string
	s1 = &s
	s = "sxy"
	println(a, *b)
	println(s, *s1)
}
