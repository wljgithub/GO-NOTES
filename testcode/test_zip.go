package main

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createZip(zipName string) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{"readme.md", "hello"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		check(err)
		_, err = f.Write([]byte(file.Body))
		check(err)
	}
	err := w.Close()
	check(err)

	file, err := os.OpenFile(zipName, os.O_RDWR|os.O_CREATE, 0777)
	defer file.Close()
	buf.WriteTo(file)

}

func toZip() {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme", "hello world in go"},
		{"main.go", "package main\nfunc main(){fmt.Println()}"}}
	for _, file := range files {

		f, err := writer.Create(file.Name)
		check(err)
		_, err = f.Write([]byte(file.Body))
		check(err)
	}
	err := writer.Close()
	check(err)
	f, err := os.OpenFile("1.zip", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	check(err)

	buf.WriteTo(f)
}
func main() {
	// toZip()
	zipName := "myzip.zip"
	createZip(zipName)

}
