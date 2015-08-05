package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/previousnext/cache/backend"
	_ "github.com/previousnext/cache/backend/folder"
	"github.com/previousnext/cache/config"
)

func main() {
	// Load the configuration file.
	f, err := config.LoadConfig(".cache.yml")
	if err != nil {
		Exit("Cannot load config file: .cache.yml")
	}

	// Load the caching backend.
	b, err := backend.New("folder")
	if err != nil {
		panic("Cannot find cache backend.")
	}

	app := cli.NewApp()
	app.Name = "Cache"
	app.Usage = "A generic caching system."
	app.Commands = []cli.Command{
		{
			Name:    "snapshot",
			Aliases: []string{"s"},
			Usage:   "Snapshot the current state",
			Action: func(c *cli.Context) {
				err := b.Snapshot(f)
				if err != nil {
					panic(err)
				}
			},
		},
		{
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "Restore the cached state",
			Action: func(c *cli.Context) {
				err := b.Restore(f)
				if err != nil {
					panic(err)
				}
			},
		},
		{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "Print out the state of the current cache",
			Action: func(c *cli.Context) {
				err := b.Print(f)
				if err != nil {
					panic(err)
				}
			},
		},
	}
	app.Run(os.Args)
}
