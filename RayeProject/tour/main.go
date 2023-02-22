package main

import (
	"github.com/raye17/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Exectute err: %v", err)
	}
}
