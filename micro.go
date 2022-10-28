package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"gitlab.com/repod/micro/cmd"
)

const (
	Name        = "micro"
	Version     = "0.0.01"
	Description = "go microservice tool kit."
)

func main() {
	//inital
	app := &cli.App{
		Name:            Name,
		Version:         Version,
		Description:     Description,
		Commands:        cmd.Commands(),
		HideHelp:        true,
		HideHelpCommand: true,
	}

	//run
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
