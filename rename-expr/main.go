package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	from, to := os.Args[1], os.Args[2]
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		new := strings.Replace(file.Name(), from, to, -1)
		os.Rename(file.Name(), new)
	}
}
