package main

import (
	"encoding/hex"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Cache struct {
	Generated       string
	Status          string
	CachedDir       string
	CachedExists    bool
	Directory       string `yaml:"directory"`
	DirectoryExists bool
	HashFiles       []string `yaml:"hash"`
}

// Builds a new file configuration.
func NewCache(file, dir string) ([]*Cache, error) {
	var caches []*Cache
	
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return caches, err
	}

	err = yaml.Unmarshal(f, &caches)
	if err != nil {
		return caches, err
	}

	// Loop over the build a hash for each of the caches.
	for _, c := range caches {
		// Compute the has values from the files on the disk.
		h, err := ComputeMd5(c.HashFiles)
        if err != nil {
        	c.Generated = "NULL"
        	c.Status = "Not cached"
        	continue
        } else {
	        c.Generated = hex.EncodeToString(h)
        }

        // This is the directory which the files are cached.
        c.CachedDir = dir + "/" + c.Generated + "/" + c.Directory

        // Check if the directory we want to cache to exists.
        c.CachedExists = DirExists(c.CachedDir)
        c.DirectoryExists = DirExists(c.Directory)
	}

	return caches, nil
}