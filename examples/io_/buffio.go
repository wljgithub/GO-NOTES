package main

import (
	"bufio"
	"os"
	"fmt"
)

func buffRead()  {
	reader:=bufio.NewReader(os.Stdin)

	reader.ReadLine()
	str,err:=reader.ReadString('\n')
	if err==nil{
		fmt.Println("input was:",str)
	}

}

func main() {
	//fmt.Println(string('A'),string('B'))
	buffRead()
}
