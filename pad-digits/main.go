package main

import (
	"io/ioutil"
	"log"
)

func main() {
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		newName := file.Name()
		log.Println("old:", file.Name(), "new:", newName)
	}
}
