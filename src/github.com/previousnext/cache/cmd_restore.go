package main

import (
	"os"
	"fmt"

	"github.com/mostafah/fsync"
	"github.com/alecthomas/kingpin"
)

type RestoreCommand struct {
	File      string
	Directory string
}

func configureCmdRestore(app *kingpin.Application) {
	c := &RestoreCommand{}
	r := app.Command("restore", "Restore from the cache").Action(c.Run)
	r.Flag("file", "Cache configuration file").Default(".cache.yml").OverrideDefaultFromEnvar("CACHE_FILE").StringVar(&c.File)
	r.Flag("dir", "The directory to cache into").Default("/var/tmp/cache").OverrideDefaultFromEnvar("CACHE_DIR").StringVar(&c.Directory)
}

func (c *RestoreCommand) Run(k *kingpin.ParseContext) error {
	// Load up the cache configuration,
	caches, err := NewCache(c.File, c.Directory)
	if err != nil {
		Exit(err.Error())
	}

	for _, cs := range caches {
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