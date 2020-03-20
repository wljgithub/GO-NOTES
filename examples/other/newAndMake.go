package main

import "fmt"

func main() {
	s:=make([]int,10)
	s:=new([]int)
	fmt.Printf("a slice created by make:%v\n",s)
	fmt.Printf("")
}
