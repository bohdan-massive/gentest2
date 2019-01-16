package main

import (
	"os"

	"github.com/bohdan-massive/gentest2/cmd"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "gentest2"
	commands := []cli.Command{
		cmd.ServeCommand(),
	}
	app.Commands = commands
	app.Run(os.Args)

}
