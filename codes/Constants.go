package main

import "fmt"

// import "math"

// 通过 const关键字声明一个常量
const s string = "constant"

func main() {

	//也可以省略声明的类型，go会自动判断
	const n = 500000
	fmt.Println(n)

	//一个数字常量在没有给定之前没有类型，可以通过强制转换类型
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
