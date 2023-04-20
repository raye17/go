package main

import (
	"os"
	"os/exec"
	"study/util/errors"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	errors.Checkout("cmd run failed: ", e)

}
