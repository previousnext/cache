package main

import (
	"fmt"
	"strings"

	"github.com/gosuri/uitable"
	"github.com/alecthomas/kingpin"
)

type ListCommand struct {
	Caches []*Cache
}

func configureCmdList(app *kingpin.Application, caches []*Cache) {
	c := &ListCommand{
		Caches: caches,
	}
	app.Command("list", "List all the caches and their status").Action(c.Run)
}

func (c *ListCommand) Run(k *kingpin.ParseContext) error {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("HASH", "FILES", "DIRECTORY", "STATUS")
	table.AddRow("----", "-----", "---------")
    for _, cs := range c.Caches {
    	if cs.CachedExists {
			table.AddRow(cs.Generated, strings.Join(cs.HashFiles, ","), cs.Directory, "Cached")
    	} else {
			table.AddRow(cs.Generated, strings.Join(cs.HashFiles, ","), cs.Directory, "Not Cached")
    	}
    }
    fmt.Println(table)
    return nil
}