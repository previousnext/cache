package main

import (
	"fmt"
	"strings"

	"github.com/gosuri/uitable"
	"github.com/alecthomas/kingpin"
)

type ListCommand struct {
	File      string
	Directory string
}

func configureCmdList(app *kingpin.Application) {
	c := &ListCommand{}
	l := app.Command("list", "List all the caches and their status").Action(c.Run)
	l.Flag("file", "Cache configuration file").Default(".cache.yml").OverrideDefaultFromEnvar("CACHE_FILE").StringVar(&c.File)
	l.Flag("dir", "The directory to cache into").Default("/var/tmp/cache").OverrideDefaultFromEnvar("CACHE_DIR").StringVar(&c.Directory)
}

func (c *ListCommand) Run(k *kingpin.ParseContext) error {
	// Load up the cache configuration,
	caches, err := NewCache(c.File, c.Directory)
	if err != nil {
		Exit(err.Error())
	}

	// Render in table form
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("HASH", "FILES", "DIRECTORY", "STATUS")
	table.AddRow("----", "-----", "---------")
    for _, cs := range caches {
    	if cs.CachedExists {
			table.AddRow(cs.Generated, strings.Join(cs.HashFiles, ","), cs.Directory, "Cached")
    	} else {
			table.AddRow(cs.Generated, strings.Join(cs.HashFiles, ","), cs.Directory, "Not Cached")
    	}
    }
    fmt.Println(table)
    return nil
}