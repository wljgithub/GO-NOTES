package main

import (
	"strconv"
	"fmt"
)

func check(err error){
	if err !=nil{
		panic(err)
	}
}
func main()  {
	number,err := strconv.Atoi("a")

	// panic 会终止程序的运行，抛出错误
	check(err)
	fmt.Println(number)
	fmt.Println("panic will not be break the program")
}
