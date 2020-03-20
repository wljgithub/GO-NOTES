package main

import "fmt"

/*
 递归时主要有两个条件：
	边界条件
	递归法则
 */

// 精简的几行代码就实现这种效果，好优雅.  递归
func recursive(num int)  {
	//先算边界条件
	if num>=10 {
		recursive(num/10)
	}
	fmt.Println(num%10)
}

func fibon(n int)int  {

	if n ==1{
		return 0
	}
	if n ==2{
		return 1
	}
	return fibon(n-1)+fibon(n-2)
}

func main()  {
	//recursive(987654321)
	fmt.Println(fibon(6))
}
