package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

// readFileLineByLineCusto 按行读取文件的内容
func readFileLineByLineCusto(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)

	}
	reader := bufio.NewReader(file)
	for {
		readStr, err := reader.ReadString('\n')
		log.Println(readStr)
		if err == io.EOF {
			return
		}
	}
}

// readFileBuiltInLBL 虽然也可以按行读取文件内容，但不推荐用这种方式，因为内置的readline,读取时不包含换行符
func readFileBuiltInLBL(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	for {
		raw, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal("occur err:", err)
		}
		log.Printf("%s", raw)

	}
}

func writeFile
func main() {
	readFileBuiltInLBL("./1.txt")

}
