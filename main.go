package main

import (
	"log"
	"os"

	"github.com/alash3al/a2a/handlers"
	"github.com/urfave/cli/v2"

	_ "github.com/alash3al/a2a/handlers/text"
)

func main() {
	app := &cli.App{
		Name:                 "a2a",
		Usage:                "convert anything to anything",
		EnableBashCompletion: true,
		Commands:             []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "src",
				Aliases: []string{"s"},
				Usage:   "the source filename",
			},
			&cli.StringFlag{
				Name:    "dest",
				Aliases: []string{"d"},
				Usage:   "the destenation filename",
			},
		},
	}

	for _, h := range handlers.Handlers() {
		h.Command.Action = wrapHandler(h.Action)
		app.Commands = append(app.Commands, &h.Command)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
