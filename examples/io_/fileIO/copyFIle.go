package main

import (
	"io"
	"log"
	"os"
)

func copyFile(from, to string) {
	fileFrom, err := os.Open(from)
	defer fileFrom.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileTo, err := os.Create(to)
	defer fileTo.Close()
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(fileTo, fileFrom)
}

func main() {
	from,to:="",""
	copyFile(from,to)
}
