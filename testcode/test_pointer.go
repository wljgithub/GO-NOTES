package main

import "fmt"

var p = 0
func noPointer(para *int)  {
	*para = 2
}

func main() {
	noPointer(&p)
	fmt.Println(p)
}
