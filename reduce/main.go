package main

import (
	"io/ioutil"
	"os"
)

func main() {
	dirs, _ := ioutil.ReadDir(".")
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		files, _ := ioutil.ReadDir(dir.Name())
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			os.Rename(dir.Name()+"/"+file.Name(), dir.Name()+"-"+file.Name())
		}
	}
}
