package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("provide a base path")
	}
	path := os.Args[1]

	wd, _ := filepath.Abs(path)
	lowerFiles(wd, path)
}

func lowerFiles(basepath, child string) {
	cwd := filepath.Join(basepath, child)
	log.Printf("processing dir: %s", cwd)
	children, err := ioutil.ReadDir(cwd)
	if err != nil {
		log.Panicf("error processing '%s': %v", cwd, err)
	}
	for _, c := range children {
		oldPath, _ := filepath.Abs(filepath.Join(cwd, c.Name()))
		newPath, _ := filepath.Abs(filepath.Join(cwd, strings.ToLower(c.Name())))
		err := os.Rename(oldPath, newPath)
		if err != nil {
			if !strings.Contains(err.Error(), ": file exists") {
				log.Fatalf("error renaming %s to %s: %v", oldPath, newPath, err)
			}
		}
		if c.IsDir() {
			lowerFiles(cwd, strings.ToLower(c.Name()))
		}
	}
}
