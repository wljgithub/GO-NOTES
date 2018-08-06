package main

import (
	"fmt"
	// "time"
)

// 用select实现管道读取超时机制
func channelTimeOut() {
	// timeout := make(chan bool, 1)
	// go func() {
	// 	time.Sleep(1e9)
	// 	timeout <- true
	// }()

	ch := make(chan int, 1)
	// ch <- 1
	// var v int
	select {
	// case <-timeout:
	case ch <- 1:
	}
	// fmt.Println("v value:", v)
	fmt.Println("end")
}

func main() {
	channelTimeOut()
	// ch := make(chan int, 1024)
	// for {
	// 	select {
	// 	case ch <- 0:
	// 	case ch <- 1:
	// 	}
	// 	i := <-ch
	// 	fmt.Println("Value Received:", i)
	// }

}
