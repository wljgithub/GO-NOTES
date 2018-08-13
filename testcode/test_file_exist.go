package main

import (
	"log"
	"os"
	"time"
)

func fileNotExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	log.Println(fileInfo.Name())
	return true
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func fileCreate(path string) *os.File {
	var file *os.File
	if fileNotExist(path) {
		file, err := os.Create(path)
		check(err)
		return file
	} else {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0777)
		check(err)
		return file
	}
	return file
}

func test() {

	a := 1
	for i := 0; i < 2; i++ {
		log.Println(a)
	}
}
func main() {
	path := "./1.txt"
	// log.Println(fileNotExist(path))
	for i := 0; i < 1000; i++ {

		file := fileCreate(path)
		file.WriteString("hello")
	}
	time.Sleep(2 * time.Second)
	// test()

}
