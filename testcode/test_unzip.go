package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
)

func readZip(path string) {
	// 打开一个zip格式文件
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("文件名 %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		s, err := ioutil.ReadAll(rc)
		fmt.Println(string(s), err)
	}
}
func main() {
	path := "i.zip"
	readZip(path)
}
