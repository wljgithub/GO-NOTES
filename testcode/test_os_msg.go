package main

import (
	"fmt"
	"log"
	"os"
)

func getOSEnv() {
	// 系统环境变量
	log.Println(os.Environ())
	log.Println(os.Getenv("PATH"))

	// 当前目录
	log.Println(os.Getwd())

	os.Exit(1)

}

func NavigateFile() {
	path := "./testCreate"
	// 创建一个文件，如果存在，则覆盖原文件(会清空原文件的内容)
	file, err := os.Create(path)
	log.Println(file.Name(), err)

	os.Remove(path)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true

}

func MyOpen() {
	// 打开一个文件，并将内容写入到文件末尾 ,(如果文件不存在则创建文件)
	path := "./1.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		file.WriteString(fmt.Sprintf("hello,%d", i))
		file.Write([]byte("\n"))

	}
}

func MyRead() {
	// 我希望有种机制，能记录我最后一次读取的位置，下一次文件读取从这个位置开始
	path := "./1.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	b1 := make([]byte, 5)
	file.Read(b1)
	fmt.Println(string(b1))

}
func main() {
	// fileInfo, err := os.Stat("./log.txt")
	// log.Println(fileInfo, err)
	// if os.IsExist(err) {
	// 	log.Println(fileInfo.Name())
	// }
	// path := "./log.txt"
	// if IsExist(path) {
	// 	os.RemoveAll(path)
	// }
	// NavigateFile()

	// MyOpen()
	MyRead()

}
