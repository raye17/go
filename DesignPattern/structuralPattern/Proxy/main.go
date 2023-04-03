package main

import (
	"fmt"
	"strings"
)

type Git interface {
	Clone(url string) bool
}
type Github struct{}

func (g Github) Clone(url string) bool {
	if strings.HasPrefix(url, "https") {
		fmt.Println("clone from " + url)
		return true
	}
	fmt.Println("failed to clone from " + url)
	return false
}

type Gitlab struct{}

func (g Gitlab) Clone(url string) bool {
	if strings.HasPrefix(url, "http") {
		fmt.Println("clone from " + url)
		return true
	}
	fmt.Println("failed to clone from " + url)
	return false
}

type GitBash struct {
	GitCmd Git
}

func (g GitBash) Clone(url string) bool {
	return g.GitCmd.Clone(url)
}

type Coder struct{}

func (c Coder) GetCode(url string, i int) {
	gitBash := GetGit(i)
	if gitBash.Clone(url) {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}
func GetGit(i int) Git {
	if i == 1 {
		return GitBash{GitCmd: Github{}}
	} else if i == 2 {
		return GitBash{
			GitCmd: Gitlab{},
		}
	} else {
		return nil

	}
}
func main() {
	c := Coder{}
	c.GetCode("https://github.com/raye17/go", 1)
	d := Coder{}
	d.GetCode("http.www.baidu.com", 2)
}
