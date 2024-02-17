package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alash3al/re-txt/handlers"
	"github.com/urfave/cli/v2"
)

// wrapHandler wrap our handler into cli.ActionFunc
func wrapHandler(fn handlers.HandlerFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		sources := [][]byte{}

		for _, src := range c.StringSlice("src") {
			srcBytes, err := ioutil.ReadFile(src)
			if err != nil {
				return err
			}
			sources = append(sources, srcBytes)
		}

		{
			stdin := os.Stdin
			stat, err := stdin.Stat()
			if err != nil {
				return err
			}

			if (stat.Mode() & os.ModeDevice) == 0 {
				stdinBytes, err := ioutil.ReadAll(stdin)
				if err != nil {
					return err
				}

				sources = append(sources, stdinBytes)
			}
		}

		result, err := fn(&handlers.Context{
			Context: c,
			Input:   sources,
		})

		if err != nil {
			return err
		}

		if c.String("dest") != "" {
			return ioutil.WriteFile(c.String("dest"), result, 0755)
		} else {
			fmt.Println(string(result))
		}

		return nil
	}
}
