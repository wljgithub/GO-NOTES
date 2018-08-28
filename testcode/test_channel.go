package main

import (
	"fmt"
	"time"
)

func concurrency(ch chan int, num int) {
	ch <- num
}
func main() {
	ch := make(chan int)
	sum := 0
	for i := 0; i < 10000000; i++ {
		go concurrency(ch, i)
	}
	go func() {
		for {
			i := <-ch
			sum += i
			if sum > 1000000000000 {

				fmt.Println("end", i)
				break
			}
			fmt.Println("counting")
		}
	}()
	time.Sleep(100e9)
	//	go func() {
	//		ch <- 1
	//		for i := 0; i < 10000; i++ {
	//			sum += i
	//		}
	//		time.Sleep(1e9)
	//	}()
	//	<-ch
}
