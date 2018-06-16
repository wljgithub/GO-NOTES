package main

import "fmt"

func main() {

	//int数组初始化后，默认值为0
	var a [5]int
	fmt.Println("emp:", a)

	//可通过索引访问或改变数组的值(索引值从0开始)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//内置函数len可返回数组长度
	fmt.Println("len:", len(a))

	//数组初始化和赋值可同时进行
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dlc", b)

	//也可以构造多维数组
	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d", twoD)

}
