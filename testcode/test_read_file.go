package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

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

func readFileBuiltInLBL(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	for {
		raw, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", raw)

	}
}

func readFileToString(path string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", raw)

}

// 指定读取文件多少个字节
// func readFileInSize(path string) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	reader := bufio.NewReader(file)
// 	reader, err = reader.Peek(8)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// raw := make([]byte, 1024)
// size, err := reader.Read(raw)

// if err != nil {
// 	log.Fatal(err)
// }

// log.Println("Content:", string(raw), "Size:", size)
// }
func writeFile(path string) {

	data := "name age hoby\njack 23 coding\nhuimin 21 TV"
	err := ioutil.WriteFile(path, []byte(data), 0600)
	if err != nil {
		log.Fatal(err)
	}

}

// 按行读取文件内容
func readFileCBC(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	var c1, c2, c3 string
	fmt.Println(c1)
	for {
		_, err := fmt.Fscanln(file, &c1, &c2, &c3)
		fmt.Println(c1, c2, c3)
		col1 = append(col1, c1)
		col1 = append(col2, c2)
		col1 = append(col3, c3)
		if err != nil {
			break
		}
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
func main() {
	// path := "./1.txt"
	// readFileLineByLine(path)
	// readFileBuiltInLBL(path)
	// readFileToString(path)
	// writeFile(path)
	// readFileInSize(path)
	// readFileCBC(path) -- 有点问题
	type s struct {
		name string
	}
	ss := s{}
	log.Println(ss == s{})
}
