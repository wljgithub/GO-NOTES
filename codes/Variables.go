package main

import "fmt"

func main() {

	//go 声明变量时，会自动判断值的类型，可以指定类型，也可以不指定

	//可以一次声明一个变量
	var a = "initial"
	fmt.Println(a)

	//也可以一次声明多个变量

	var b, c int = 1, 2
	fmt.Println(b, c)

	//未指定类型，go则会自行判断
	var d = true
	fmt.Println(d)

	//初始化变量时，若不赋值，则为默认值
	var e int
	fmt.Println(e)

	// := 符号代表初始化并赋值给一个变量
	f := "short"
	fmt.Println(f)

}
