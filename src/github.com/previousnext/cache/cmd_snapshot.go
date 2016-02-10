package main

import (
	"fmt"
	"os"

	"github.com/mostafah/fsync"
	"github.com/alecthomas/kingpin"
)

type SnapshotCommand struct {
	Caches []*Cache
}

func configureCmdSnapshot(app *kingpin.Application, caches []*Cache) {
	c := &SnapshotCommand{
		Caches: caches,
	}
	app.Command("snapshot", "Take a snapshot of the current directories").Action(c.Run)
}

func (c *SnapshotCommand) Run(k *kingpin.ParseContext) error {
	for _, cs := range c.Caches {
		// Check the cache directory exists.
		if !cs.DirectoryExists {
	    	fmt.Printf("Cannot find directory to snapshot: %s\n", cs.Directory)
	    	continue
		}

		// Remove the target directory.
		os.RemoveAll(cs.CachedDir)
        os.MkdirAll(cs.CachedDir, 0777)

		// Copy our cache into the new directory.
        err := fsync.Sync(cs.CachedDir, cs.Directory)
        if err != nil {
            fmt.Printf("Failed to snapshot %s from directory: %s\n", cs.Directory, err.Error())
        } else {
            fmt.Printf("Successfully snapshotted: %s\n", cs.Directory)
        }
	}

	return nil
}