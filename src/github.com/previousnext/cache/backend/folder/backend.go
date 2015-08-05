package folder

import (
	"fmt"
	"os"

	"github.com/mostafah/fsync"
	"github.com/olekukonko/tablewriter"

	backend "github.com/previousnext/cache/backend"
	"github.com/previousnext/cache/config"
)

var (
	folderDir = "/tmp/cache"
)

type FolderBackend struct{}

func init() {
	backend.Register("folder", &FolderBackend{})
}

func (o *FolderBackend) Snapshot(c []config.Config) error {
	// Remove the directory since we are going to rebuild it.
	os.RemoveAll(folderDir)

	for _, v := range c {
		// This is the hash of the file which we use to determine the diretory
		// which it is stored in.
		h, e := v.Hash()
		if e != nil {
			Output("Snapshot", "Hash does not exist for "+v.HashFileFlat())
			continue
		}

		// Loop over the folders.
		for _, f := range v.Restore {
			p := folderDir + "/" + h + "/" + f

			// Blow away and recreate the snaphost directory.
			os.MkdirAll(p, 0777)

			// Copy the directory contents to this new directory.
			err := fsync.Sync(p, f)
			if err != nil {
				Output("Snapshot", err.Error())
			} else {
				Output("Snapshot", v.HashFileFlat()+" - "+f+" to "+p)
			}
		}
	}
	return nil
}

func (o *FolderBackend) Restore(c []config.Config) error {
	for _, v := range c {
		// This is the hash of the file which we use to determine the diretory
		// which it is stored in.
		h, e := v.Hash()
		if e != nil {
			Output("Restore", "Hash does not exist for "+v.HashFileFlat())
			continue
		}

		// Loop over the folders.
		for _, f := range v.Restore {
			p := folderDir + "/" + h + "/" + f

			// Blow away and recreate the target directory.
			os.RemoveAll(f)
			os.MkdirAll(f, 0777)

			// Copy the directory contents to this new directory.
			err := fsync.Sync(f, p)
			if err != nil {
				Output("Restore", err.Error())
			} else {
				Output("Restore", v.HashFileFlat()+" - "+p+" to "+f)
			}
		}
	}
	return nil
}

// Helper function to print to the screen in a consistent way.
func (o *FolderBackend) Print(c []config.Config) error {
	var data [][]string
	for _, i := range c {
		// This is a row in the table to be printed.
		h, e := i.Hash()
		if e != nil {
			h = "NULL"
		}

		// Check if this file/folders have a cache setup.
		cd := "Not cached"
		_, err := os.Stat(folderDir + "/" + h)
		if err == nil {
			cd = "Cached"
		}

		n := []string{
			h,
			i.HashFileFlat(),
			i.RestoreFlat(),
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
