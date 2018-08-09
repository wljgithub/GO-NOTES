package main

import (
	"io/ioutil"
	"log"
	"os"
)

func iterOverDir(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		log.Printf(" file name :%s --- size :%v", fileInfo.Name(), fileInfo.Size())
	}

}
func main() {
	// err := ioutil.WriteFile("./log.txt", []byte("hellom,wrold"), 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	os.Open("./1.txt")
}
