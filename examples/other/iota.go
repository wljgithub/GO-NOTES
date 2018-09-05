package main

import (
	"fmt"
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const NextIota = iota

func main() {

	// iota 初始为0，每次调用会加一，
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)

	// 直到下一个const关键字出现，就会重置为0
	fmt.Println("iota:", NextIota)
}
