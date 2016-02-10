package main

import (
	"fmt"
	"os"

	"github.com/mostafah/fsync"
	"github.com/alecthomas/kingpin"
)

type SnapshotCommand struct {
	File      string
	Directory string
}

func configureCmdSnapshot(app *kingpin.Application) {
	c := &SnapshotCommand{}
	s := app.Command("snapshot", "Take a snapshot of the current directories").Action(c.Run)
	s.Flag("file", "Cache configuration file").Default(".cache.yml").OverrideDefaultFromEnvar("CACHE_FILE").StringVar(&c.File)
	s.Flag("dir", "The directory to cache into").Default("/var/tmp/cache").OverrideDefaultFromEnvar("CACHE_DIR").StringVar(&c.Directory)
}

func (c *SnapshotCommand) Run(k *kingpin.ParseContext) error {
	// Load up the cache configuration,
	caches, err := NewCache(c.File, c.Directory)
	if err != nil {
		Exit(err.Error())
	}

	for _, cs := range caches {
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