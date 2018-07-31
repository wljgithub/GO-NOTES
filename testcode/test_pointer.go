package main

import "fmt"

var p = 0
func Pointer(para *int)  {
	*para = 2
}

func NoPointer(para int)  {
	para = 5
}

func main() {
	Pointer(&p)
	fmt.Printf("使用用指针后，p的值为:%v",p)

	NoPointer(p)
	fmt.Printf("不使用指针，p的值为：%v",p)
}
