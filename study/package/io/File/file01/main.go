package main

import (
	"fmt"
	"os"
)

func main() {
	proverbs := []string{
		"hello,world\n",
		"my name is raye\n",
		"The programming language I'm learning is go\n",
		"Don't give up\n",
	}
	file, err := os.Create("./proverbs.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	for _, p := range proverbs {
		n, err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(err, "file write failed!")
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data!")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
}
