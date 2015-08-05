package folder

import (
	"fmt"
	"os"
	"strings"

	"github.com/mostafah/fsync"
	"github.com/olekukonko/tablewriter"

	backend "github.com/previousnext/cache/backend"
	"github.com/previousnext/cache/config"
)

type FolderBackend struct{}

func init() {
	backend.Register("folder", &FolderBackend{})
}

func (o *FolderBackend) Snapshot(c []config.Config) error {
	d := "/tmp/cache"

	// Remove the directory since we are going to rebuild it.
	os.RemoveAll(d)

	for _, v := range c {
		// This is the hash of the file which we use to determine the diretory
		// which it is stored in.
		h, e := v.Hash()
		if e != nil {
			Output("Snapshot", "Hash does not exist for "+v.HashFile)
			continue
		}

		// Loop over the folders.
		for _, f := range v.Restore {
			p := d + "/" + h + "/" + f

			// Blow away and recreate the snaphost directory.
			os.MkdirAll(p, 0644)

			// Copy the directory contents to this new directory.
			err := fsync.Sync(p, f)
			if err != nil {
				Output("Snapshot", err.Error())
			} else {
				Output("Snapshot", v.HashFile+" - "+f+" to "+p)
			}
		}
	}
	return nil
}

func (o *FolderBackend) Restore(c []config.Config) error {
	d := "/tmp/cache"

	for _, v := range c {
		// This is the hash of the file which we use to determine the diretory
		// which it is stored in.
		h, e := v.Hash()
		if e != nil {
			Output("Restore", "Hash does not exist for "+v.HashFile)
			continue
		}

		// Loop over the folders.
		for _, f := range v.Restore {
			p := d + "/" + h + "/" + f

			// Blow away and recreate the target directory.
			os.RemoveAll(f)
			os.MkdirAll(f, 0644)

			// Copy the directory contents to this new directory.
			err := fsync.Sync(f, p)
			if err != nil {
				Output("Restore", err.Error())
			} else {
				Output("Restore", v.HashFile+" - "+p+" to "+f)
			}
		}
	}
	return nil
}

// Helper function to print to the screen in a consistent way.
func (o *FolderBackend) Print(c []config.Config) error {
	d := "/tmp/cache"

	var data [][]string
	for _, i := range c {
		// This is a row in the table to be printed.
		h, e := i.Hash()
		if e != nil {
			h = "NULL"
		}

		// Check if this file/folders have a cache setup.
		cd := "Not cached"
		_, err := os.Stat(d + "/" + h)
		if err == nil {
			cd = "Cached"
		}

		n := []string{
			h,
			i.HashFile,
			strings.Join(i.Restore, ","),
			cd,
		}
		data = append(data, n)
	}

	// Print out the data in a structured format.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hash", "File", "Restore", "Status"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
	return nil
}

func Output(t string, m string) {
	fmt.Println(t + ": " + m)
}
