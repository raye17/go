package main

import (
	"os"
	"os/exec"
	"study/util"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	util.Checkout("cmd run failed: ", e)

}
