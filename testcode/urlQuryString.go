package main

import (
	"fmt"
	)

func main()  {
	//fmt.Println(url.QueryUnescape("name%3D%E5%B7%AB%E5%87%8C%E5%81%A5"))
	//userInfo:=url.UserPassword("jack","*****")
	//fmt.Println(userInfo.Username())
	//fmt.Println(userInfo.Password())
	var c = [...]int{1,2,3,3:0}
	fmt.Println(c)

	var a = []int{1, 2, 3} // a 是一个数组
	var b = a                // b 是指向数组的指针

	fmt.Println(a[0], a[1])   // 打印数组的前2个元素
	a[0]=0
	fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似

	//for i, v := range b {     // 通过数组指针迭代数组的元素
	//	fmt.Println(i, v)
	//}


}