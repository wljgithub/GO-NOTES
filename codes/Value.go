package main

import "fmt"

func main() {

	//go 的数据类型包括字符串、整型、浮点型、布尔类型

	//字符串可以通过 + 号来拼接
	fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0+3.0=", 7.0+3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
