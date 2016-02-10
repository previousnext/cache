package main

import (
	"os"

	"github.com/alecthomas/kingpin"
)

const (
	version  = "0.0.1"
	file     = ".cache.yml"
	cacheDir = "/tmp/cache"
)

func main() {
	kingpin.Version(version)

	app := kingpin.New("Cache", "A generic cache control system")
	cache, err := NewCache(file, cacheDir)
	if err != nil {
		Exit(err.Error())
	}

	configureCmdList(app, cache)
	configureCmdRestore(app, cache)
	configureCmdSnapshot(app, cache)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
