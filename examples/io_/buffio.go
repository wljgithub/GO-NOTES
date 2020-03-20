package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// buffRead 从标准输入读取用户输入，并输出
func buffRead() {
	reader := bufio.NewReader(os.Stdin)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if str == "quit" {
			break
		}
		fmt.Println("input was:", str)
	}
}

func readAndPrint() {
	for {
		io.Copy(os.Stdout, os.Stdin)
	}
}

//func sayHello() {
//	for {
//	output:=bufio.NewWriter(os.Stdout)
//		n,err:=output.ReadFrom(os.Stdin
//		fmt.Fprint(os.Stdout, "hello")
//	}
//}
func main() {
	buffRead()
	readAndPrint()
}
