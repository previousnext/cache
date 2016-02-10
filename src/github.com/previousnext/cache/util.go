package main

import (
	"fmt"
	"crypto/md5"
    "io"
    "os"
)

func Exit(m string) {
	fmt.Println(m)
	os.Exit(1)
}

func DirExists(d string) bool {
    _, err := os.Stat(d)
    if err != nil {
    	return false
    }
	return true
}

func ComputeMd5(files []string) ([]byte, error) {
    var result []byte

    hash := md5.New()
    for _, f := range files {
        file, err := os.Open(f)
        if err != nil {
                return result, err
        }
        defer file.Close()

        _, err = io.Copy(hash, file)
        if err != nil {
                return result, err
        }
    }

    return hash.Sum(result), nil
}
