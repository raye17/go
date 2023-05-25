package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func main() {
	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:     "ls-l",
				Usage:    "列出目录",
				Category: "directory",
				Action: func(c *cli.Context) error {
					fmt.Println("args[0]:", c.Args().Get(0))
					return nil
				},
			},
			{
				Name:     "mkdir-m",
				Aliases:  []string{"m"},
				Category: "directory",
				Usage:    "创建文件夹",
				Action: func(c *cli.Context) error {
					fmt.Println("mkdir...")
					return nil
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "test for test",
				Subcommands: []*cli.Command{
					{
						Name:  "sub-test01",
						Usage: "show sub-test01",
						Action: func(c *cli.Context) error {
							fmt.Println("this is sub-test01")
							return nil
						},
					},
					{
						Name: "sub-test02",
						Action: func(c *cli.Context) error {
							fmt.Println("this is sub-test02")
							return nil
						},
					},
				},
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		fmt.Println("command error:", err.Error())
	}
}
