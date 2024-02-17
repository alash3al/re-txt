package text

import (
	"github.com/alash3al/re-txt/handlers"
	"github.com/urfave/cli/v2"
)

func init() {
	prettyFlag := &cli.BoolFlag{
		Name:    "pretty",
		Aliases: []string{"p"},
		Value:   true,
	}

	handlers.Handle(&handlers.Handler{
		Command: cli.Command{
			Name:     "yaml2json",
			Usage:    "convert from yaml to json",
			Aliases:  []string{"y2j"},
			Category: "Text Converters",
			Flags: []cli.Flag{
				prettyFlag,
			},
		},

		Action: handlers.HandlerFunc(yaml2json),
	})

	handlers.Handle(&handlers.Handler{
		Command: cli.Command{
			Name:     "json2yaml",
			Usage:    "convert from json to yaml",
			Aliases:  []string{"j2y"},
			Category: "Text Converters",
		},

		Action: handlers.HandlerFunc(json2yaml),
	})

	handlers.Handle(&handlers.Handler{
		Command: cli.Command{
			Name:     "toml2json",
			Usage:    "convert from toml to json",
			Aliases:  []string{"t2j"},
			Category: "Text Converters",
			Flags: []cli.Flag{
				prettyFlag,
			},
		},

		Action: handlers.HandlerFunc(toml2json),
	})

	handlers.Handle(&handlers.Handler{
		Command: cli.Command{
			Name:     "json2toml",
			Usage:    "convert from json to toml",
			Aliases:  []string{"j2t"},
			Category: "Text Converters",
		},

		Action: handlers.HandlerFunc(json2toml),
	})

	handlers.Handle(&handlers.Handler{
		Command: cli.Command{
			Name:     "hcl2json",
			Usage:    "convert from hcl to json",
			Aliases:  []string{"h2j"},
			Category: "Text Converters",
			Flags: []cli.Flag{
				prettyFlag,
			},
		},

		Action: handlers.HandlerFunc(hcl2json),
	})

	handlers.Handle(&handlers.Handler{
		Command: cli.Command{
			Name:     "csv2json",
			Usage:    "convert from csv to json",
			Aliases:  []string{"c2j"},
			Category: "Text Converters",
			Flags: []cli.Flag{
				prettyFlag,
				&cli.StringFlag{
					Name:     "header",
					Required: true,
				},
				&cli.BoolFlag{
					Name:     "contains-header",
					Required: true,
					Value:    false,
				},
				&cli.StringFlag{
					Name:    "separator",
					Aliases: []string{"s"},
					Value:   ",",
				},
				&cli.StringFlag{
					Name:    "comment",
					Aliases: []string{"c"},
					Value:   "#",
				},
				&cli.BoolFlag{
					Name:  "trim-leading-space",
					Value: true,
				},
				&cli.BoolFlag{
					Name:  "lazy-quotes",
					Value: true,
				},
			},
		},

		Action: handlers.HandlerFunc(csv2json),
	})
}
