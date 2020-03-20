package main

import (
	"log"
	"runtime"
)

func write(ch chan int) {
	log.Println("counting")
	ch <- 1

}

var x int

// 不要用共享内存来通信，要用通信来共享内存
func main() {
	chs := make([]chan int, 10)
	for i, _ := range chs {
		chs[i] = make(chan int)
		go write(chs[i])
	}

	for _, ch := range chs {
		x = <-ch
		// log.Println(x)
	}
	runtime.Gosched()

	for x >= 5 {
		log.Println(x)

	}

}
