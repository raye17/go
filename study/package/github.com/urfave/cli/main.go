package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

type Command struct {
}

func (c *Command) Test(cli *cli.Context) {
	uid := cli.Int("uid")
	username := cli.String("username")
	fmt.Println(uid, username)
}
func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Commands = []cli.Command{
		{
			Name:   "test",
			Usage:  "test --uid=x --username=y",
			Action: (&Command{}).Test,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "uid",
					Usage: "--uid",
				},
				cli.StringFlag{
					Name:  "username",
					Usage: "--username",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("command error:", err.Error())
	}
}
