package main

import (
	"github.com/RushikeshMarkad16/Library-Managemant/app"
	"github.com/RushikeshMarkad16/Library-Managemant/config"
	"github.com/RushikeshMarkad16/Library-Managemant/server"
	"github.com/urfave/cli"
)

func main() {
	config.Load()
	app.Init()
	defer app.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "Library Management App"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
	}
}
