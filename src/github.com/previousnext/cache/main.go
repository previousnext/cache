package main

import (
	"os"

	"github.com/alecthomas/kingpin"

	"github.com/previousnext/cache/backend"
	_ "github.com/previousnext/cache/backend/folder"
	"github.com/previousnext/cache/config"
)

var (
	cliApp      = kingpin.New("cache", "A generic caching system")
	cliFile     = cliApp.Flag("file", "Path to configuration file").Default(".cache.yml").String()
	cliBackend  = cliApp.Flag("backend", "The name of the backend").Default("folder").String()
	cliSnapshot = cliApp.Command("snapshot", "Snapshot the current directories")
	cliRestore  = cliApp.Command("restore", "Restore to the current directories")
	cliStatus   = cliApp.Command("status", "Print out the current status of what is cached")
)

func main() {
	switch kingpin.MustParse(cliApp.Parse(os.Args[1:])) {
	case cliSnapshot.FullCommand():
		f := GetConfig(*cliFile)
		b := GetBackend(*cliBackend)
		err := b.Snapshot(f)
		if err != nil {
			panic(err)
		}
	case cliRestore.FullCommand():
		f := GetConfig(*cliFile)
		b := GetBackend(*cliBackend)
		err := b.Restore(f)
		if err != nil {
			panic(err)
		}
	case cliStatus.FullCommand():
		f := GetConfig(*cliFile)
		b := GetBackend(*cliBackend)
		err := b.Print(f)
		if err != nil {
			panic(err)
		}
	}
}

func GetConfig(f string) []config.Config {
	c, err := config.LoadConfig(f)
	if err != nil {
		Exit("Cannot load config file: .cache.yml")
	}
	return c
}

func GetBackend(n string) backend.Backend {
	b, err := backend.New(n)
	if err != nil {
		panic("Cannot find cache backend.")
	}
	return b
}
