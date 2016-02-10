package main

import (
	"os"
	"fmt"

	"github.com/mostafah/fsync"
	"github.com/alecthomas/kingpin"
)

type RestoreCommand struct {
	Caches []*Cache
}

func configureCmdRestore(app *kingpin.Application, caches []*Cache) {
	c := &RestoreCommand{
		Caches: caches,
	}
	app.Command("restore", "Restore from the cache").Action(c.Run)
}

func (c *RestoreCommand) Run(k *kingpin.ParseContext) error {
	for _, cs := range c.Caches {
		// Check the cache directory exists.
		if !cs.CachedExists {
	    	fmt.Printf("Cannot find cache for: %s\n", cs.Directory)
	    	continue
		}

		// Remove the target directory.
		os.RemoveAll(cs.Directory)
        os.MkdirAll(cs.Directory, 0777)

		// Copy our cache into the new directory.
        err := fsync.Sync(cs.Directory, cs.CachedDir)
        if err != nil {
            fmt.Printf("Failed to restore %s from cache: %s\n", cs.Directory, err.Error())
        } else {
            fmt.Printf("Successfully restored: %s\n", cs.Directory)
        }
	}

	return nil
}