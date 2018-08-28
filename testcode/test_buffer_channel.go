package main

import (
	"fmt"
)

var ch = make(chan int, 100)

func bufferChannel(i int) {
	ch <- 1
	fmt.Println(i)
	<-ch
}
func main() {
	for i := 0; i < 10000; i++ {
		go bufferChannel(i)
	}
	fmt.Scanln()
}
