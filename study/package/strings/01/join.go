package main

import (
	"fmt"
	"strings"
)

func main() {
	versions := []string{"v1", "v2", "v3"}
	fmt.Println(String(versions))
	v := verbs{"create", "list", "delete"}
	v2 := []string{"put", "patch"}
	v1 := verbs(v2)
	s1 := v1.String()
	s := v.String()
	fmt.Println(s)
	fmt.Println(s1)
}
func String(v []string) string {
	return strings.Join(v, ",")
}

type verbs []string

func (v verbs) String() string {
	return fmt.Sprintf("%v", []string(v))
}
