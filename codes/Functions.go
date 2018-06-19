package main

import "fmt"

// 传入两个整型，返回他们的和
func plus(a int, b int) int {

	// go中不会自动返回值，需要明确指明返回值
	return a + b
}

// 如果要传入的类型一致，可一次性声明
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// 调用函数
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
