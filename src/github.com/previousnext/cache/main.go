package main

import (
	"os"

	"github.com/alecthomas/kingpin"
)

func main() {
	app := kingpin.New("Cache", "A generic cache control system")

	configureCmdList(app)
	configureCmdRestore(app)
	configureCmdSnapshot(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
